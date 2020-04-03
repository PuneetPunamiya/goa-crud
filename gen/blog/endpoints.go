// Code generated by goa v3.1.1, DO NOT EDIT.
//
// blog endpoints
//
// Command:
// $ goa gen crud/design

package blog

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "blog" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	List   goa.Endpoint
	Remove goa.Endpoint
	Update goa.Endpoint
	Add    goa.Endpoint
	Show   goa.Endpoint
}

// NewEndpoints wraps the methods of the "blog" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
		List:   NewListEndpoint(s),
		Remove: NewRemoveEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Add:    NewAddEndpoint(s),
		Show:   NewShowEndpoint(s),
	}
}

// Use applies the given middleware to all the "blog" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.List = m(e.List)
	e.Remove = m(e.Remove)
	e.Update = m(e.Update)
	e.Add = m(e.Add)
	e.Show = m(e.Show)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "blog".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Blog)
		return s.Create(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "blog".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.List(ctx)
	}
}

// NewRemoveEndpoint returns an endpoint function that calls the method
// "remove" of service "blog".
func NewRemoveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemovePayload)
		return nil, s.Remove(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "blog".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
		return nil, s.Update(ctx, p)
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "blog".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*NewComment)
		return s.Add(ctx, p)
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "blog".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Blog)
		return s.Show(ctx, p)
	}
}
