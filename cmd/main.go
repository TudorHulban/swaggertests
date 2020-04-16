// Package main is used for testing Swagger file generation.
//
// Documentation for API.
//
package main

import (
	"github.com/TudorHulban/swaggertests/pkg/handlers"
	ugo "github.com/savsgio/atreugo/v11"
)

func main() {
	config := ugo.Config{
		Addr:    "0.0.0.0:8000",
		GetOnly: true,
	}

	server := ugo.New(config)

	server.GET("/r1", handlers.HandlerRoute1)
	server.GET("/r2", handlers.HandlerRoute2)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
