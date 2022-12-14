// Code generated by goa v3.8.1, DO NOT EDIT.
//
// chatter HTTP server
//
// Command:
// $ goa gen github.com/sevein/oneof/design -o .

package server

import (
	"context"
	"net/http"

	chatter "github.com/sevein/oneof/gen/chatter"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the chatter service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	Subscribe http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the chatter service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *chatter.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer *ConnConfigurer,
) *Server {
	if configurer == nil {
		configurer = &ConnConfigurer{}
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Subscribe", "GET", "/subscribe"},
		},
		Subscribe: NewSubscribeHandler(e.Subscribe, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.SubscribeFn),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "chatter" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Subscribe = m(s.Subscribe)
}

// Mount configures the mux to serve the chatter endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountSubscribeHandler(mux, h.Subscribe)
}

// Mount configures the mux to serve the chatter endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountSubscribeHandler configures the mux to serve the "chatter" service
// "subscribe" endpoint.
func MountSubscribeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/subscribe", f)
}

// NewSubscribeHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "subscribe" endpoint.
func NewSubscribeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		encodeError = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		var err error
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &chatter.SubscribeEndpointInput{
			Stream: &SubscribeServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
