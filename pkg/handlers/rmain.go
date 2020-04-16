package handlers

import (
	ugo "github.com/savsgio/atreugo/v11"
)

// swagger:route GET /r1 ThisIsOperationID
// This is a summary of the route.
//
// This is description line 1.
// This is description line 2.
//
// responses:
//	200: OK

// HandlerRoute1 Returns OK.
func HandlerRoute1(ctx *ugo.RequestCtx) error {
	return ctx.HTTPResponse("OK", 200)
}

// swagger:route GET /r2 HTTPResponse
// Route r2 does R2. This is a summary of the route.
//
// responses:
//	200: OK

// HandlerRoute2 Returns OK.
func HandlerRoute2(ctx *ugo.RequestCtx) error {
	return ctx.HTTPResponse("OK", 200)
}

// swagger :response Resp
type swaggerResp struct {
}
