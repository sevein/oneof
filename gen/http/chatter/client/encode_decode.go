// Code generated by goa v3.7.13, DO NOT EDIT.
//
// chatter HTTP client encoders and decoders
//
// Command:
// $ goa-v3.7.13 gen github.com/sevein/oneof/design -o .

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	chatter "github.com/sevein/oneof/gen/chatter"
	chatterviews "github.com/sevein/oneof/gen/chatter/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildSubscribeRequest instantiates a HTTP request object with method and
// path set to call the "chatter" service "subscribe" endpoint
func (c *Client) BuildSubscribeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: SubscribeChatterPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chatter", "subscribe", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeSubscribeResponse returns a decoder for responses returned by the
// chatter subscribe endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeSubscribeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SubscribeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chatter", "subscribe", err)
			}
			p := NewSubscribeOneofEventOK(&body)
			view := "default"
			vres := &chatterviews.OneofEvent{Projected: p, View: view}
			if err = chatterviews.ValidateOneofEvent(vres); err != nil {
				return nil, goahttp.ErrValidationError("chatter", "subscribe", err)
			}
			res := chatter.NewOneofEvent(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chatter", "subscribe", resp.StatusCode, string(body))
		}
	}
}
