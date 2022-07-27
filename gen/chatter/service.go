// Code generated by goa v3.7.13, DO NOT EDIT.
//
// chatter service
//
// Command:
// $ goa-v3.7.13 gen github.com/sevein/oneof/design -o .

package chatter

import (
	"context"

	chatterviews "github.com/sevein/oneof/gen/chatter/views"
)

// The chatter service implements a simple client and server chat.
type Service interface {
	// Subscribe to events sent when new chat messages are added.
	Subscribe(context.Context, SubscribeServerStream) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "chatter"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"subscribe"}

// SubscribeServerStream is the interface a "subscribe" endpoint server stream
// must satisfy.
type SubscribeServerStream interface {
	// Send streams instances of "OneofEvent".
	Send(*OneofEvent) error
	// Close closes the stream.
	Close() error
}

// SubscribeClientStream is the interface a "subscribe" endpoint client stream
// must satisfy.
type SubscribeClientStream interface {
	// Recv reads instances of "OneofEvent" from the stream.
	Recv() (*OneofEvent, error)
}

// OneofEvent is the result type of the chatter service subscribe method.
type OneofEvent struct {
	Payload interface {
		payloadVal()
	}
}

type OneofFoobarEvent struct {
	Message *string
}

type OneofPingEvent struct {
	Message *string
}

func (*OneofFoobarEvent) payloadVal() {}
func (*OneofPingEvent) payloadVal()   {}

// NewOneofEvent initializes result type OneofEvent from viewed result type
// OneofEvent.
func NewOneofEvent(vres *chatterviews.OneofEvent) *OneofEvent {
	return newOneofEvent(vres.Projected)
}

// NewViewedOneofEvent initializes viewed result type OneofEvent from result
// type OneofEvent using the given view.
func NewViewedOneofEvent(res *OneofEvent, view string) *chatterviews.OneofEvent {
	p := newOneofEventView(res)
	return &chatterviews.OneofEvent{Projected: p, View: "default"}
}

// newOneofEvent converts projected type OneofEvent to service type OneofEvent.
func newOneofEvent(vres *chatterviews.OneofEventView) *OneofEvent {
	res := &OneofEvent{}
	if vres.Payload != nil {
		switch actual := vres.Payload.(type) {
		case *OneofPingEvent:
			val := &OneofFoobarEvent{
				Message: actual.Message,
			}
			res.Payload = Payload{Value: val}
		case *OneofFoobarEvent:
			val := &OneofPingEvent{
				Message: actual.Message,
			}
			res.Payload = Payload{Value: val}
		}
	}
	return res
}

// newOneofEventView projects result type OneofEvent to projected type
// OneofEventView using the "default" view.
func newOneofEventView(res *OneofEvent) *chatterviews.OneofEventView {
	vres := &chatterviews.OneofEventView{}
	if res.Payload != nil {
		switch actual := res.Payload.(type) {
		case *chatterviews.OneofPingEvent:
			val := &chatterviews.OneofFoobarEvent{
				Message: actual.Message,
			}
			vres.Payload = chatterviews.Payload{Value: val}
		case *chatterviews.OneofFoobarEvent:
			val := &chatterviews.OneofPingEvent{
				Message: actual.Message,
			}
			vres.Payload = chatterviews.Payload{Value: val}
		}
	}
	return vres
}

// transformChatterviewsOneofFoobarEventToOneofFoobarEvent builds a value of
// type *OneofFoobarEvent from a value of type *chatterviews.OneofFoobarEvent.
func transformChatterviewsOneofFoobarEventToOneofFoobarEvent(v *chatterviews.OneofFoobarEvent) *OneofFoobarEvent {
	if v == nil {
		return nil
	}
	res := &OneofFoobarEvent{
		Message: v.Message,
	}

	return res
}

// transformChatterviewsOneofPingEventToOneofPingEvent builds a value of type
// *OneofPingEvent from a value of type *chatterviews.OneofPingEvent.
func transformChatterviewsOneofPingEventToOneofPingEvent(v *chatterviews.OneofPingEvent) *OneofPingEvent {
	if v == nil {
		return nil
	}
	res := &OneofPingEvent{
		Message: v.Message,
	}

	return res
}

// transformOneofFoobarEventToChatterviewsOneofFoobarEvent builds a value of
// type *chatterviews.OneofFoobarEvent from a value of type *OneofFoobarEvent.
func transformOneofFoobarEventToChatterviewsOneofFoobarEvent(v *OneofFoobarEvent) *chatterviews.OneofFoobarEvent {
	if v == nil {
		return nil
	}
	res := &chatterviews.OneofFoobarEvent{
		Message: v.Message,
	}

	return res
}

// transformOneofPingEventToChatterviewsOneofPingEvent builds a value of type
// *chatterviews.OneofPingEvent from a value of type *OneofPingEvent.
func transformOneofPingEventToChatterviewsOneofPingEvent(v *OneofPingEvent) *chatterviews.OneofPingEvent {
	if v == nil {
		return nil
	}
	res := &chatterviews.OneofPingEvent{
		Message: v.Message,
	}

	return res
}