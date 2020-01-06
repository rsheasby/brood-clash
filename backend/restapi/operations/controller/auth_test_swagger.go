// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AuthTestHandlerFunc turns a function with the right signature into a auth test handler
type AuthTestHandlerFunc func(AuthTestParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AuthTestHandlerFunc) Handle(params AuthTestParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AuthTestHandler interface for that can handle valid auth test params
type AuthTestHandler interface {
	Handle(AuthTestParams, interface{}) middleware.Responder
}

// NewAuthTest creates a new http.Handler for the auth test operation
func NewAuthTest(ctx *middleware.Context, handler AuthTestHandler) *AuthTest {
	return &AuthTest{Context: ctx, Handler: handler}
}

/*AuthTest swagger:route GET /controller/authTest controller authTest

Test the authorization code

*/
type AuthTest struct {
	Context *middleware.Context
	Handler AuthTestHandler
}

func (o *AuthTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAuthTestParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
