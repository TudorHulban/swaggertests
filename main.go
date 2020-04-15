package main

import (
	ugo "github.com/savsgio/atreugo"
)

func handlerRoute(ctx *ugo.RequestCtx) error {
	return ctx.HTTPResponse("OK", 200)
}

func main() {
	config := ugo.Config{
		Addr:    "0.0.0.0:8000",
		GetOnly: true,
	}

	server := ugo.New(config)

	server.GET("/", handlerRoute)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
