package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sevein/oneof/gen/chatter"
	chattersvc "github.com/sevein/oneof/gen/chatter"
	chattersvr "github.com/sevein/oneof/gen/http/chatter/server"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[chatterapi] ", log.Ltime)
	}

	// Initialize the services.
	var (
		chatterSvc chatter.Service
	)
	{
		chatterSvc = NewChatter(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		chatterEndpoints *chatter.Endpoints
	)
	{
		chatterEndpoints = chatter.NewEndpoints(chatterSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		addr := "http://localhost:80"
		u, err := url.Parse(addr)
		if err != nil {
			logger.Fatalf("invalid URL %#v: %s\n", addr, err)
		}
		if *secureF {
			u.Scheme = "https"
		}
		if *domainF != "" {
			u.Host = *domainF
		}
		if *httpPortF != "" {
			h, _, err := net.SplitHostPort(u.Host)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
			}
			u.Host = net.JoinHostPort(h, *httpPortF)
		} else if u.Port() == "" {
			u.Host = net.JoinHostPort(u.Host, "80")
		}
		handleHTTPServer(ctx, u, chatterEndpoints, &wg, errc, logger, *dbgF)

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, chatterEndpoints *chatter.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {
	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		chatterServer *chattersvr.Server
	)
	{
		eh := errorHandler(logger)
		upgrader := &websocket.Upgrader{}
		chatterServer = chattersvr.New(chatterEndpoints, mux, dec, enc, eh, nil, upgrader, nil)
		if debug {
			servers := goahttp.Servers{
				chatterServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	chattersvr.Mount(mux, chatterServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range chatterServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			logger.Printf("failed to shutdown: %v", err)
		}
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}

// chatter service example implementation.
// The example methods log the requests and return zero values.
type chatterSvc struct {
	logger *log.Logger
}

// NewChatter returns the chatter service implementation.
func NewChatter(logger *log.Logger) chattersvc.Service {
	return &chatterSvc{
		logger: logger,
	}
}

func (s *chatterSvc) Subscribe(ctx context.Context, stream chattersvc.SubscribeServerStream) (err error) {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	done := false
	for {
		select {
		case <-ctx.Done():
			done = true
		case <-ticker.C:
			message := "ping"
			stream.Send(&chattersvc.OneofEvent{
				Payload: &chattersvc.OneofPingEvent{
					Message: &message,
				},
			})
		}
		if done {
			break
		}
	}

	return stream.Close()
}
