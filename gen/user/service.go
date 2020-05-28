// Code generated by goa v3.1.1, DO NOT EDIT.
//
// user service
//
// Command:
// $ goa gen github.com/sm43/goa-crud/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The user service gives user details.
type Service interface {
	// Add a new blog
	Create(context.Context, *CreatePayload) (err error)
	// List all the users
	List(context.Context) (res []*User, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "user"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"create", "list"}

// CreatePayload is the payload type of the user service create method.
type CreatePayload struct {
	// Adding a new user
	User *User
	// Access github token
	Auth string
}

// A User describes a user retrieved by the storage service.
type User struct {
	// ID is the unique id of the user
	ID *uint
	// Name of user
	Name string
	// Age of user
	Age uint
	// Class of user
	Class string
}

// MakeDbError builds a goa.ServiceError from an error.
func MakeDbError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "db_error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInvalidToken builds a goa.ServiceError from an error.
func MakeInvalidToken(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "invalid-token",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}