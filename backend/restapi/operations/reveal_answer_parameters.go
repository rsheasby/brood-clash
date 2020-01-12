// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRevealAnswerParams creates a new RevealAnswerParams object
// no default values defined in spec.
func NewRevealAnswerParams() RevealAnswerParams {

	return RevealAnswerParams{}
}

// RevealAnswerParams contains all the bound params for the reveal answer operation
// typically these are obtained from a http.Request
//
// swagger:parameters revealAnswer
type RevealAnswerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	AnswerID int64
	/*
	  Required: true
	  In: path
	*/
	QuestionID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRevealAnswerParams() beforehand.
func (o *RevealAnswerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAnswerID, rhkAnswerID, _ := route.Params.GetOK("answerId")
	if err := o.bindAnswerID(rAnswerID, rhkAnswerID, route.Formats); err != nil {
		res = append(res, err)
	}

	rQuestionID, rhkQuestionID, _ := route.Params.GetOK("questionId")
	if err := o.bindQuestionID(rQuestionID, rhkQuestionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAnswerID binds and validates parameter AnswerID from path.
func (o *RevealAnswerParams) bindAnswerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("answerId", "path", "int64", raw)
	}
	o.AnswerID = value

	return nil
}

// bindQuestionID binds and validates parameter QuestionID from path.
func (o *RevealAnswerParams) bindQuestionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("questionId", "path", "int64", raw)
	}
	o.QuestionID = value

	return nil
}