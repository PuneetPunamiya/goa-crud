// Code generated by goa v3.1.1, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/sm43/goa-crud/design

package server

import (
	user "github.com/sm43/goa-crud/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "user" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Adding a new user
	User *UserRequestBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// ListResponseBody is the type of the "user" service "list" endpoint HTTP
// response body.
type ListResponseBody []*UserResponse

// CreateDbErrorResponseBody is the type of the "user" service "create"
// endpoint HTTP response body for the "db_error" error.
type CreateDbErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateInvalidTokenResponseBody is the type of the "user" service "create"
// endpoint HTTP response body for the "invalid-token" error.
type CreateInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ListDbErrorResponseBody is the type of the "user" service "list" endpoint
// HTTP response body for the "db_error" error.
type ListDbErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	// ID is the unique id of the user
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of user
	Name string `form:"name" json:"name" xml:"name"`
	// Age of user
	Age uint `form:"age" json:"age" xml:"age"`
	// Class of user
	Class string `form:"class" json:"class" xml:"class"`
}

// UserRequestBody is used to define fields on request body types.
type UserRequestBody struct {
	// ID is the unique id of the user
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Age of user
	Age *uint `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
	// Class of user
	Class *string `form:"class,omitempty" json:"class,omitempty" xml:"class,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "user" service.
func NewListResponseBody(res []*user.User) ListResponseBody {
	body := make([]*UserResponse, len(res))
	for i, val := range res {
		body[i] = marshalUserUserToUserResponse(val)
	}
	return body
}

// NewCreateDbErrorResponseBody builds the HTTP response body from the result
// of the "create" endpoint of the "user" service.
func NewCreateDbErrorResponseBody(res *goa.ServiceError) *CreateDbErrorResponseBody {
	body := &CreateDbErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateInvalidTokenResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "user" service.
func NewCreateInvalidTokenResponseBody(res *goa.ServiceError) *CreateInvalidTokenResponseBody {
	body := &CreateInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListDbErrorResponseBody builds the HTTP response body from the result of
// the "list" endpoint of the "user" service.
func NewListDbErrorResponseBody(res *goa.ServiceError) *ListDbErrorResponseBody {
	body := &ListDbErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreatePayload builds a user service create endpoint payload.
func NewCreatePayload(body *CreateRequestBody, auth string) *user.CreatePayload {
	v := &user.CreatePayload{}
	v.User = unmarshalUserRequestBodyToUserUser(body.User)
	v.Auth = auth

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserRequestBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUserRequestBody runs the validations defined on UserRequestBody
func ValidateUserRequestBody(body *UserRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Age == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("age", "body"))
	}
	if body.Class == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("class", "body"))
	}
	return
}