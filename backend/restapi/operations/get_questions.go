// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetQuestionsHandlerFunc turns a function with the right signature into a get questions handler
type GetQuestionsHandlerFunc func(GetQuestionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetQuestionsHandlerFunc) Handle(params GetQuestionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetQuestionsHandler interface for that can handle valid get questions params
type GetQuestionsHandler interface {
	Handle(GetQuestionsParams, interface{}) middleware.Responder
}

// NewGetQuestions creates a new http.Handler for the get questions operation
func NewGetQuestions(ctx *middleware.Context, handler GetQuestionsHandler) *GetQuestions {
	return &GetQuestions{Context: ctx, Handler: handler}
}

/*GetQuestions swagger:route GET /questions getQuestions

Get all questions

*/
type GetQuestions struct {
	Context *middleware.Context
	Handler GetQuestionsHandler
}

func (o *GetQuestions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetQuestionsParams()

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
