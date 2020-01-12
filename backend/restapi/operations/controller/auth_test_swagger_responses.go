// Code generated by go-swagger; DO NOT EDIT.

package controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AuthTestOKCode is the HTTP code returned for type AuthTestOK
const AuthTestOKCode int = 200

/*AuthTestOK OK

swagger:response authTestOK
*/
type AuthTestOK struct {
}

// NewAuthTestOK creates AuthTestOK with default headers values
func NewAuthTestOK() *AuthTestOK {

	return &AuthTestOK{}
}

// WriteResponse to the client
func (o *AuthTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
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