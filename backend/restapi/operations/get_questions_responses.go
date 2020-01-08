// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "backend/models"
)

// GetQuestionsOKCode is the HTTP code returned for type GetQuestionsOK
const GetQuestionsOKCode int = 200

/*GetQuestionsOK OK

swagger:response getQuestionsOK
*/
type GetQuestionsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Question `json:"body,omitempty"`
}

// NewGetQuestionsOK creates GetQuestionsOK with default headers values
func NewGetQuestionsOK() *GetQuestionsOK {

	return &GetQuestionsOK{}
}

// WithPayload adds the payload to the get questions o k response
func (o *GetQuestionsOK) WithPayload(payload []*models.Question) *GetQuestionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get questions o k response
func (o *GetQuestionsOK) SetPayload(payload []*models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetQuestionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Question, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetQuestionsUnauthorizedCode is the HTTP code returned for type GetQuestionsUnauthorized
const GetQuestionsUnauthorizedCode int = 401

/*GetQuestionsUnauthorized Unauthorized

swagger:response getQuestionsUnauthorized
*/
type GetQuestionsUnauthorized struct {
}

// NewGetQuestionsUnauthorized creates GetQuestionsUnauthorized with default headers values
func NewGetQuestionsUnauthorized() *GetQuestionsUnauthorized {

	return &GetQuestionsUnauthorized{}
}

// WriteResponse to the client
func (o *GetQuestionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

/*GetQuestionsDefault Unknown error

swagger:response getQuestionsDefault
*/
type GetQuestionsDefault struct {
	_statusCode int
}

// NewGetQuestionsDefault creates GetQuestionsDefault with default headers values
func NewGetQuestionsDefault(code int) *GetQuestionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetQuestionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get questions default response
func (o *GetQuestionsDefault) WithStatusCode(code int) *GetQuestionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get questions default response
func (o *GetQuestionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *GetQuestionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
