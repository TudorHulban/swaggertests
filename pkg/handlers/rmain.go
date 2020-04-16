package handlers

import (
	ugo "github.com/savsgio/atreugo/v11"
)

// swagger:route GET /r1 HTTPResponse
// Route r1 does R1. This is a summary of the route.
//
// responses:
//	200: OK

func HandlerRoute1(ctx *ugo.RequestCtx) error {
	return ctx.HTTPResponse("OK", 200)
}

// swagger:route GET /r2 HTTPResponse
// Route r2 does R2. This is a summary of the route.
//
// responses:
//	200: OK

func HandlerRoute2(ctx *ugo.RequestCtx) error {
	return ctx.HTTPResponse("OK", 200)
}
