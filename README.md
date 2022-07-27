Build issues in Goa-generated code when using `OneOf` in a result type used
in `StreamingResult`:

```
# github.com/sevein/oneof/gen/chatter
gen/chatter/service.go:84:8: impossible type switch case: *OneofPingEvent
	vres.Payload (variable of type interface{payloadVal()}) cannot have dynamic type *OneofPingEvent (missing payloadVal method)
gen/chatter/service.go:88:18: undefined: Payload
gen/chatter/service.go:89:8: impossible type switch case: *OneofFoobarEvent
	vres.Payload (variable of type interface{payloadVal()}) cannot have dynamic type *OneofFoobarEvent (missing payloadVal method)
gen/chatter/service.go:93:18: undefined: Payload
gen/chatter/service.go:105:22: undefined: views.OneofPingEvent
gen/chatter/service.go:106:25: undefined: views.OneofFoobarEvent
gen/chatter/service.go:122:78: undefined: views.OneofFoobarEvent
gen/chatter/service.go:135:74: undefined: views.OneofPingEvent
gen/chatter/service.go:148:97: undefined: views.OneofFoobarEvent
gen/chatter/service.go:161:91: undefined: views.OneofPingEvent
gen/chatter/service.go:106:25: too many errors
```
