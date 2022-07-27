// Code generated by goa v3.7.13, DO NOT EDIT.
//
// chatter endpoints
//
// Command:
// $ goa-v3.7.13 gen github.com/sevein/oneof/design -o .

package chatter

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "chatter" service endpoints.
type Endpoints struct {
	Subscribe goa.Endpoint
}

// SubscribeEndpointInput holds both the payload and the server stream of the
// "subscribe" method.
type SubscribeEndpointInput struct {
	// Stream is the server stream used by the "subscribe" method to send data.
	Stream SubscribeServerStream
}

// NewEndpoints wraps the methods of the "chatter" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Subscribe: NewSubscribeEndpoint(s),
	}
}

// Use applies the given middleware to all the "chatter" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Subscribe = m(e.Subscribe)
}

// NewSubscribeEndpoint returns an endpoint function that calls the method
// "subscribe" of service "chatter".
func NewSubscribeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*SubscribeEndpointInput)
		return nil, s.Subscribe(ctx, ep.Stream)
	}
}