.DEFAULT_GOAL := goa

goa:
	go run -mod=readonly goa.design/goa/v3/cmd/goa gen github.com/sevein/oneof/design -o .
