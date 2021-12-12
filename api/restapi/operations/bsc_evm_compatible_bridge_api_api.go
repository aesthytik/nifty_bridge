// Code generated by go-swagger; DO NOT EDIT.

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

	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_1155_swap_pairs"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_1155_swaps"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swap_pairs"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swaps"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/svc_info"
)

// NewBscEvmCompatibleBridgeAPIAPI creates a new BscEvmCompatibleBridgeAPI instance
func NewBscEvmCompatibleBridgeAPIAPI(spec *loads.Document) *BscEvmCompatibleBridgeAPIAPI {
	return &BscEvmCompatibleBridgeAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		Erc1155SwapPairsGetErc1155SwapPairsHandler: erc_1155_swap_pairs.GetErc1155SwapPairsHandlerFunc(func(params erc_1155_swap_pairs.GetErc1155SwapPairsParams) middleware.Responder {
			return middleware.NotImplemented("operation erc_1155_swap_pairs.GetErc1155SwapPairs has not yet been implemented")
		}),
		Erc1155SwapsGetErc1155SwapsHandler: erc_1155_swaps.GetErc1155SwapsHandlerFunc(func(params erc_1155_swaps.GetErc1155SwapsParams) middleware.Responder {
			return middleware.NotImplemented("operation erc_1155_swaps.GetErc1155Swaps has not yet been implemented")
		}),
		Erc721SwapPairsGetErc721SwapPairsHandler: erc_721_swap_pairs.GetErc721SwapPairsHandlerFunc(func(params erc_721_swap_pairs.GetErc721SwapPairsParams) middleware.Responder {
			return middleware.NotImplemented("operation erc_721_swap_pairs.GetErc721SwapPairs has not yet been implemented")
		}),
		Erc721SwapsGetErc721SwapsHandler: erc_721_swaps.GetErc721SwapsHandlerFunc(func(params erc_721_swaps.GetErc721SwapsParams) middleware.Responder {
			return middleware.NotImplemented("operation erc_721_swaps.GetErc721Swaps has not yet been implemented")
		}),
		SvcInfoGetInfoHandler: svc_info.GetInfoHandlerFunc(func(params svc_info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation svc_info.GetInfo has not yet been implemented")
		}),
	}
}

/*BscEvmCompatibleBridgeAPIAPI The BSC <-> EVM Compatible Swap API: provide swap service between BSC and EVM Compatible, which is based on https://github.com/binance-chain/eth-swap-ap */
type BscEvmCompatibleBridgeAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

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

	// Erc1155SwapPairsGetErc1155SwapPairsHandler sets the operation handler for the get erc1155 swap pairs operation
	Erc1155SwapPairsGetErc1155SwapPairsHandler erc_1155_swap_pairs.GetErc1155SwapPairsHandler
	// Erc1155SwapsGetErc1155SwapsHandler sets the operation handler for the get erc1155 swaps operation
	Erc1155SwapsGetErc1155SwapsHandler erc_1155_swaps.GetErc1155SwapsHandler
	// Erc721SwapPairsGetErc721SwapPairsHandler sets the operation handler for the get erc721 swap pairs operation
	Erc721SwapPairsGetErc721SwapPairsHandler erc_721_swap_pairs.GetErc721SwapPairsHandler
	// Erc721SwapsGetErc721SwapsHandler sets the operation handler for the get erc721 swaps operation
	Erc721SwapsGetErc721SwapsHandler erc_721_swaps.GetErc721SwapsHandler
	// SvcInfoGetInfoHandler sets the operation handler for the get info operation
	SvcInfoGetInfoHandler svc_info.GetInfoHandler

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

// UseRedoc for documentation at /docs
func (o *BscEvmCompatibleBridgeAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *BscEvmCompatibleBridgeAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *BscEvmCompatibleBridgeAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BscEvmCompatibleBridgeAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BscEvmCompatibleBridgeAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BscEvmCompatibleBridgeAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BscEvmCompatibleBridgeAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BscEvmCompatibleBridgeAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BscEvmCompatibleBridgeAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BscEvmCompatibleBridgeAPIAPI
func (o *BscEvmCompatibleBridgeAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.Erc1155SwapPairsGetErc1155SwapPairsHandler == nil {
		unregistered = append(unregistered, "erc_1155_swap_pairs.GetErc1155SwapPairsHandler")
	}
	if o.Erc1155SwapsGetErc1155SwapsHandler == nil {
		unregistered = append(unregistered, "erc_1155_swaps.GetErc1155SwapsHandler")
	}
	if o.Erc721SwapPairsGetErc721SwapPairsHandler == nil {
		unregistered = append(unregistered, "erc_721_swap_pairs.GetErc721SwapPairsHandler")
	}
	if o.Erc721SwapsGetErc721SwapsHandler == nil {
		unregistered = append(unregistered, "erc_721_swaps.GetErc721SwapsHandler")
	}
	if o.SvcInfoGetInfoHandler == nil {
		unregistered = append(unregistered, "svc_info.GetInfoHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BscEvmCompatibleBridgeAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BscEvmCompatibleBridgeAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *BscEvmCompatibleBridgeAPIAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *BscEvmCompatibleBridgeAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *BscEvmCompatibleBridgeAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *BscEvmCompatibleBridgeAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the bsc evm compatible bridge API API
func (o *BscEvmCompatibleBridgeAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BscEvmCompatibleBridgeAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/erc-1155-swap-pairs"] = erc_1155_swap_pairs.NewGetErc1155SwapPairs(o.context, o.Erc1155SwapPairsGetErc1155SwapPairsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/erc-1155-swaps"] = erc_1155_swaps.NewGetErc1155Swaps(o.context, o.Erc1155SwapsGetErc1155SwapsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/erc-721-swap-pairs"] = erc_721_swap_pairs.NewGetErc721SwapPairs(o.context, o.Erc721SwapPairsGetErc721SwapPairsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/erc-721-swaps"] = erc_721_swaps.NewGetErc721Swaps(o.context, o.Erc721SwapsGetErc721SwapsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/info"] = svc_info.NewGetInfo(o.context, o.SvcInfoGetInfoHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BscEvmCompatibleBridgeAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *BscEvmCompatibleBridgeAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *BscEvmCompatibleBridgeAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *BscEvmCompatibleBridgeAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *BscEvmCompatibleBridgeAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}