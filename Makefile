.DEFAULT_GOAL := goa

goa:
	@goa-v3.7.13 gen github.com/sevein/oneof/design -o .
