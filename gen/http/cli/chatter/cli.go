// Code generated by goa v3.7.13, DO NOT EDIT.
//
// chatter HTTP client CLI support package
//
// Command:
// $ goa-v3.7.13 gen github.com/sevein/oneof/design -o .

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	chatterc "github.com/sevein/oneof/gen/http/chatter/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `chatter subscribe
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` chatter subscribe` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	chatterConfigurer *chatterc.ConnConfigurer,
) (goa.Endpoint, interface{}, error) {
	var (
		chatterFlags = flag.NewFlagSet("chatter", flag.ContinueOnError)

		chatterSubscribeFlags = flag.NewFlagSet("subscribe", flag.ExitOnError)
	)
	chatterFlags.Usage = chatterUsage
	chatterSubscribeFlags.Usage = chatterSubscribeUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "chatter":
			svcf = chatterFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "chatter":
			switch epn {
			case "subscribe":
				epf = chatterSubscribeFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "chatter":
			c := chatterc.NewClient(scheme, host, doer, enc, dec, restore, dialer, chatterConfigurer)
			switch epn {
			case "subscribe":
				endpoint = c.Subscribe()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// chatterUsage displays the usage of the chatter command and its subcommands.
func chatterUsage() {
	fmt.Fprintf(os.Stderr, `The chatter service implements a simple client and server chat.
Usage:
    %[1]s [globalflags] chatter COMMAND [flags]

COMMAND:
    subscribe: Subscribe to events sent when new chat messages are added.

Additional help:
    %[1]s chatter COMMAND --help
`, os.Args[0])
}
func chatterSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chatter subscribe

Subscribe to events sent when new chat messages are added.

Example:
    %[1]s chatter subscribe
`, os.Args[0])
}
