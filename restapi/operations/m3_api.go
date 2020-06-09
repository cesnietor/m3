// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/minio/m3/models"
	"github.com/minio/m3/restapi/operations/admin_api"
	"github.com/minio/m3/restapi/operations/user_api"
)

// NewM3API creates a new M3 instance
func NewM3API(spec *loads.Document) *M3API {
	return &M3API{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		AdminAPICreateTenantHandler: admin_api.CreateTenantHandlerFunc(func(params admin_api.CreateTenantParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_api.CreateTenant has not yet been implemented")
		}),
		AdminAPIDeleteTenantHandler: admin_api.DeleteTenantHandlerFunc(func(params admin_api.DeleteTenantParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_api.DeleteTenant has not yet been implemented")
		}),
		AdminAPIListStorageClassesHandler: admin_api.ListStorageClassesHandlerFunc(func(params admin_api.ListStorageClassesParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_api.ListStorageClasses has not yet been implemented")
		}),
		AdminAPIListTenantsHandler: admin_api.ListTenantsHandlerFunc(func(params admin_api.ListTenantsParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_api.ListTenants has not yet been implemented")
		}),
		UserAPILoginHandler: user_api.LoginHandlerFunc(func(params user_api.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user_api.Login has not yet been implemented")
		}),
		UserAPILoginDetailHandler: user_api.LoginDetailHandlerFunc(func(params user_api.LoginDetailParams) middleware.Responder {
			return middleware.NotImplemented("operation user_api.LoginDetail has not yet been implemented")
		}),
		UserAPILoginOauth2AuthHandler: user_api.LoginOauth2AuthHandlerFunc(func(params user_api.LoginOauth2AuthParams) middleware.Responder {
			return middleware.NotImplemented("operation user_api.LoginOauth2Auth has not yet been implemented")
		}),
		UserAPILogoutHandler: user_api.LogoutHandlerFunc(func(params user_api.LogoutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user_api.Logout has not yet been implemented")
		}),
		UserAPISessionCheckHandler: user_api.SessionCheckHandlerFunc(func(params user_api.SessionCheckParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation user_api.SessionCheck has not yet been implemented")
		}),
		AdminAPITenantInfoHandler: admin_api.TenantInfoHandlerFunc(func(params admin_api.TenantInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation admin_api.TenantInfo has not yet been implemented")
		}),

		KeyAuth: func(token string, scopes []string) (*models.Principal, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (key) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*M3API the m3 API */
type M3API struct {
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

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// KeyAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	KeyAuth func(string, []string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// AdminAPICreateTenantHandler sets the operation handler for the create tenant operation
	AdminAPICreateTenantHandler admin_api.CreateTenantHandler
	// AdminAPIDeleteTenantHandler sets the operation handler for the delete tenant operation
	AdminAPIDeleteTenantHandler admin_api.DeleteTenantHandler
	// AdminAPIListStorageClassesHandler sets the operation handler for the list storage classes operation
	AdminAPIListStorageClassesHandler admin_api.ListStorageClassesHandler
	// AdminAPIListTenantsHandler sets the operation handler for the list tenants operation
	AdminAPIListTenantsHandler admin_api.ListTenantsHandler
	// UserAPILoginHandler sets the operation handler for the login operation
	UserAPILoginHandler user_api.LoginHandler
	// UserAPILoginDetailHandler sets the operation handler for the login detail operation
	UserAPILoginDetailHandler user_api.LoginDetailHandler
	// UserAPILoginOauth2AuthHandler sets the operation handler for the login oauth2 auth operation
	UserAPILoginOauth2AuthHandler user_api.LoginOauth2AuthHandler
	// UserAPILogoutHandler sets the operation handler for the logout operation
	UserAPILogoutHandler user_api.LogoutHandler
	// UserAPISessionCheckHandler sets the operation handler for the session check operation
	UserAPISessionCheckHandler user_api.SessionCheckHandler
	// AdminAPITenantInfoHandler sets the operation handler for the tenant info operation
	AdminAPITenantInfoHandler admin_api.TenantInfoHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *M3API) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *M3API) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *M3API) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *M3API) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *M3API) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *M3API) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *M3API) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the M3API
func (o *M3API) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeyAuth == nil {
		unregistered = append(unregistered, "KeyAuth")
	}

	if o.AdminAPICreateTenantHandler == nil {
		unregistered = append(unregistered, "admin_api.CreateTenantHandler")
	}
	if o.AdminAPIDeleteTenantHandler == nil {
		unregistered = append(unregistered, "admin_api.DeleteTenantHandler")
	}
	if o.AdminAPIListStorageClassesHandler == nil {
		unregistered = append(unregistered, "admin_api.ListStorageClassesHandler")
	}
	if o.AdminAPIListTenantsHandler == nil {
		unregistered = append(unregistered, "admin_api.ListTenantsHandler")
	}
	if o.UserAPILoginHandler == nil {
		unregistered = append(unregistered, "user_api.LoginHandler")
	}
	if o.UserAPILoginDetailHandler == nil {
		unregistered = append(unregistered, "user_api.LoginDetailHandler")
	}
	if o.UserAPILoginOauth2AuthHandler == nil {
		unregistered = append(unregistered, "user_api.LoginOauth2AuthHandler")
	}
	if o.UserAPILogoutHandler == nil {
		unregistered = append(unregistered, "user_api.LogoutHandler")
	}
	if o.UserAPISessionCheckHandler == nil {
		unregistered = append(unregistered, "user_api.SessionCheckHandler")
	}
	if o.AdminAPITenantInfoHandler == nil {
		unregistered = append(unregistered, "admin_api.TenantInfoHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *M3API) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *M3API) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "key":
			result[name] = o.BearerAuthenticator(name, func(token string, scopes []string) (interface{}, error) {
				return o.KeyAuth(token, scopes)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *M3API) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *M3API) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
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

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *M3API) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
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
func (o *M3API) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the m3 API
func (o *M3API) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *M3API) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tenants"] = admin_api.NewCreateTenant(o.context, o.AdminAPICreateTenantHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/tenants/{name}"] = admin_api.NewDeleteTenant(o.context, o.AdminAPIDeleteTenantHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/storage-classes"] = admin_api.NewListStorageClasses(o.context, o.AdminAPIListStorageClassesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tenants"] = admin_api.NewListTenants(o.context, o.AdminAPIListTenantsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = user_api.NewLogin(o.context, o.UserAPILoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/login"] = user_api.NewLoginDetail(o.context, o.UserAPILoginDetailHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login/oauth2/auth"] = user_api.NewLoginOauth2Auth(o.context, o.UserAPILoginOauth2AuthHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/logout"] = user_api.NewLogout(o.context, o.UserAPILogoutHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/session"] = user_api.NewSessionCheck(o.context, o.UserAPISessionCheckHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tenants/{name}"] = admin_api.NewTenantInfo(o.context, o.AdminAPITenantInfoHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *M3API) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *M3API) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *M3API) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *M3API) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *M3API) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
