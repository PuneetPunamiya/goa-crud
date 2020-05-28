// Code generated by goa v3.1.1, DO NOT EDIT.
//
// user endpoints
//
// Command:
// $ goa gen github.com/sm43/goa-crud/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "user" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	List   goa.Endpoint
}

// NewEndpoints wraps the methods of the "user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
		List:   NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "user" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.List = m(e.List)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "user".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatePayload)
		return nil, s.Create(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "user".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.List(ctx)
	}
}
