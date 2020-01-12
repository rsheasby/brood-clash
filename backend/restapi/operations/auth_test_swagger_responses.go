// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AuthTestNoContentCode is the HTTP code returned for type AuthTestNoContent
const AuthTestNoContentCode int = 204

/*AuthTestNoContent No Content

swagger:response authTestNoContent
*/
type AuthTestNoContent struct {
}

// NewAuthTestNoContent creates AuthTestNoContent with default headers values
func NewAuthTestNoContent() *AuthTestNoContent {

	return &AuthTestNoContent{}
}

// WriteResponse to the client
func (o *AuthTestNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// AuthTestUnauthorizedCode is the HTTP code returned for type AuthTestUnauthorized
const AuthTestUnauthorizedCode int = 401

/*AuthTestUnauthorized Unauthorized

swagger:response authTestUnauthorized
*/
type AuthTestUnauthorized struct {
}

// NewAuthTestUnauthorized creates AuthTestUnauthorized with default headers values
func NewAuthTestUnauthorized() *AuthTestUnauthorized {

	return &AuthTestUnauthorized{}
}

// WriteResponse to the client
func (o *AuthTestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

/*AuthTestDefault Unknown error

swagger:response authTestDefault
*/
type AuthTestDefault struct {
	_statusCode int
}

// NewAuthTestDefault creates AuthTestDefault with default headers values
func NewAuthTestDefault(code int) *AuthTestDefault {
	if code <= 0 {
		code = 500
	}

	return &AuthTestDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the auth test default response
func (o *AuthTestDefault) WithStatusCode(code int) *AuthTestDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the auth test default response
func (o *AuthTestDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *AuthTestDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}