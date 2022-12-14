// Code generated by goa v3.8.1, DO NOT EDIT.
//
// chatter WebSocket client streaming
//
// Command:
// $ goa gen github.com/sevein/oneof/design -o .

package client

import (
	"io"

	"github.com/gorilla/websocket"
	chatter "github.com/sevein/oneof/gen/chatter"
	chatterviews "github.com/sevein/oneof/gen/chatter/views"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "chatter" service.
type ConnConfigurer struct {
	SubscribeFn goahttp.ConnConfigureFunc
}

// SubscribeClientStream implements the chatter.SubscribeClientStream interface.
type SubscribeClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "chatter" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		SubscribeFn: fn,
	}
}

// Recv reads instances of "chatter.OneofEvent" from the "subscribe" endpoint
// websocket connection.
func (s *SubscribeClientStream) Recv() (*chatter.OneofEvent, error) {
	var (
		rv   *chatter.OneofEvent
		body SubscribeResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewSubscribeOneofEventOK(&body)
	vres := &chatterviews.OneofEvent{res, "default"}
	if err := chatterviews.ValidateOneofEvent(vres); err != nil {
		return rv, goahttp.ErrValidationError("chatter", "subscribe", err)
	}
	return chatter.NewOneofEvent(vres), nil
}
