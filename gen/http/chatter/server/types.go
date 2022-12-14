// Code generated by goa v3.8.1, DO NOT EDIT.
//
// chatter HTTP server types
//
// Command:
// $ goa gen github.com/sevein/oneof/design -o .

package server

import (
	"encoding/json"

	chatterviews "github.com/sevein/oneof/gen/chatter/views"
)

// SubscribeResponseBody is the type of the "chatter" service "subscribe"
// endpoint HTTP response body.
type SubscribeResponseBody struct {
	Payload *struct {
		// Union type name, one of:
		// - "ping_event"
		// - "foobar_event"
		Type string `form:"Type" json:"Type" xml:"Type"`
		// JSON formatted union value
		Value string `form:"Value" json:"Value" xml:"Value"`
	} `form:"payload,omitempty" json:"payload,omitempty" xml:"payload,omitempty"`
}

// NewSubscribeResponseBody builds the HTTP response body from the result of
// the "subscribe" endpoint of the "chatter" service.
func NewSubscribeResponseBody(res *chatterviews.OneofEventView) *SubscribeResponseBody {
	body := &SubscribeResponseBody{}
	if res.Payload != nil {
		js, _ := json.Marshal(res.Payload)
		var name string
		switch res.Payload.(type) {
		case *chatterviews.OneofPingEventView:
			name = "ping_event"
		case *chatterviews.OneofFoobarEventView:
			name = "foobar_event"
		}
		body.Payload = &struct {
			// Union type name, one of:
			// - "ping_event"
			// - "foobar_event"
			Type string `form:"Type" json:"Type" xml:"Type"`
			// JSON formatted union value
			Value string `form:"Value" json:"Value" xml:"Value"`
		}{
			Type:  name,
			Value: string(js),
		}
	}
	return body
}
