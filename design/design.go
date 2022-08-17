package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("chatter", func() {
	Title("Chatter service describing the streaming features of goa v2.")
})

var _ = Service("chatter", func() {
	Description("The chatter service implements a simple client and server chat.")

	Method("subscribe", func() {
		Description("Subscribe to events sent when new chat messages are added.")
		StreamingResult(Event, func() {
			View("default")
		})
		HTTP(func() {
			GET("/subscribe")
			Response(StatusOK)
		})
	})
})

var Event = ResultType("application/vnd.oneof.event", func() {
	Attributes(func() {
		OneOf("payload", func() {
			Attribute(
				"ping_event",
				PingEvent,
				func() { View("default") },
			)
			Attribute(
				"foobar_event",
				FoobarEvent,
				func() { View("default") },
			)
		})
	})

	View("default", func() {
		Attribute("payload")
	})
})

var PingEvent = ResultType("application/vnd.oneof.ping-event", func() {
	Attributes(func() {
		Attribute("message", String)
	})

	View("default", func() {
		Attribute("message")
	})
})

var FoobarEvent = ResultType("application/vnd.oneof.foobar-event", func() {
	Attributes(func() {
		Attribute("message", String)
		Attribute("item", StoredItem, func() { View("default") })
	})

	View("default", func() {
		Attribute("message")
		Attribute("item")
	})
})

var StoredItem = ResultType("application/vnd.oneof.stored-item", func() {
	Description("StoredItem describes an item retrieved by the service.")
	Attributes(func() {
		Attribute("id", UInt, "Identifier of item")
		Attribute("name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
	Required("id", "name")
})
