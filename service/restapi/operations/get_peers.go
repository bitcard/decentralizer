package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetPeersHandlerFunc turns a function with the right signature into a get peers handler
type GetPeersHandlerFunc func(GetPeersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPeersHandlerFunc) Handle(params GetPeersParams) middleware.Responder {
	return fn(params)
}

// GetPeersHandler interface for that can handle valid get peers params
type GetPeersHandler interface {
	Handle(GetPeersParams) middleware.Responder
}

// NewGetPeers creates a new http.Handler for the get peers operation
func NewGetPeers(ctx *middleware.Context, handler GetPeersHandler) *GetPeers {
	return &GetPeers{Context: ctx, Handler: handler}
}

/*GetPeers swagger:route GET /v1/peers/{identifier} getPeers

Gets ips that are also looking on this identifier

*/
type GetPeers struct {
	Context *middleware.Context
	Handler GetPeersHandler
}

func (o *GetPeers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetPeersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
