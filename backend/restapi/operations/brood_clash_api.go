// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"backend/restapi/operations/controller"
)

// NewBroodClashAPI creates a new BroodClash instance
func NewBroodClashAPI(spec *loads.Document) *BroodClashAPI {
	return &BroodClashAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ControllerAuthTestHandler: controller.AuthTestHandlerFunc(func(params controller.AuthTestParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ControllerAuthTest has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		APIKeyAuthAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (ApiKeyAuth) Authorization from header param [Authorization] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*BroodClashAPI Backend API for Brood Clash */
type BroodClashAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// APIKeyAuthAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	APIKeyAuthAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ControllerAuthTestHandler sets the operation handler for the auth test operation
	ControllerAuthTestHandler controller.AuthTestHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *BroodClashAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BroodClashAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BroodClashAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BroodClashAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BroodClashAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BroodClashAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BroodClashAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BroodClashAPI
func (o *BroodClashAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuthAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.ControllerAuthTestHandler == nil {
		unregistered = append(unregistered, "controller.AuthTestHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BroodClashAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BroodClashAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {

		case "ApiKeyAuth":

			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuthAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *BroodClashAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *BroodClashAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *BroodClashAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BroodClashAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the brood clash API
func (o *BroodClashAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BroodClashAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/controller/authTest"] = controller.NewAuthTest(o.context, o.ControllerAuthTestHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BroodClashAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BroodClashAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BroodClashAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BroodClashAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
