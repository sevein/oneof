// Code generated by goa v3.7.14, DO NOT EDIT.
//
// chatter views
//
// Command:
// $ goa gen github.com/sevein/oneof/design -o .

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// OneofEvent is the viewed result type that is projected based on a view.
type OneofEvent struct {
	// Type to project
	Projected *OneofEventView
	// View to render
	View string
}

// OneofEventView is a type that runs validations on a projected type.
type OneofEventView struct {
	Payload interface {
		payloadVal()
	}
}

var (
	// OneofEventMap is a map indexing the attribute names of OneofEvent by view
	// name.
	OneofEventMap = map[string][]string{
		"default": {
			"payload",
		},
	}
)

// ValidateOneofEvent runs the validations defined on the viewed result type
// OneofEvent.
func ValidateOneofEvent(result *OneofEvent) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateOneofEventView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateOneofEventView runs the validations defined on OneofEventView using
// the "default" view.
func ValidateOneofEventView(result *OneofEventView) (err error) {

	return
}
