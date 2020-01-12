// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetQuestionHandlerFunc turns a function with the right signature into a get question handler
type GetQuestionHandlerFunc func(GetQuestionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetQuestionHandlerFunc) Handle(params GetQuestionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetQuestionHandler interface for that can handle valid get question params
type GetQuestionHandler interface {
	Handle(GetQuestionParams, interface{}) middleware.Responder
}

// NewGetQuestion creates a new http.Handler for the get question operation
func NewGetQuestion(ctx *middleware.Context, handler GetQuestionHandler) *GetQuestion {
	return &GetQuestion{Context: ctx, Handler: handler}
}

/*GetQuestion swagger:route GET /questions/{id} getQuestion

Get a question by ID

*/
type GetQuestion struct {
	Context *middleware.Context
	Handler GetQuestionHandler
}

func (o *GetQuestion) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetQuestionParams()

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