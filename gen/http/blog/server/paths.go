// Code generated by goa v3.1.1, DO NOT EDIT.
//
// HTTP request path constructors for the blog service.
//
// Command:
// $ goa gen crud/design

package server

import (
	"fmt"
)

// CreateBlogPath returns the URL path to the blog service create HTTP endpoint.
func CreateBlogPath() string {
	return "/"
}

// ListBlogPath returns the URL path to the blog service list HTTP endpoint.
func ListBlogPath() string {
	return "/"
}

// RemoveBlogPath returns the URL path to the blog service remove HTTP endpoint.
func RemoveBlogPath(id uint32) string {
	return fmt.Sprintf("/%v", id)
}

// UpdateBlogPath returns the URL path to the blog service update HTTP endpoint.
func UpdateBlogPath(id uint32) string {
	return fmt.Sprintf("/%v", id)
}

// AddBlogPath returns the URL path to the blog service add HTTP endpoint.
func AddBlogPath(id uint32) string {
	return fmt.Sprintf("/%v/comments", id)
}

// ShowBlogPath returns the URL path to the blog service show HTTP endpoint.
func ShowBlogPath(id uint32) string {
	return fmt.Sprintf("/%v", id)
}

// OauthBlogPath returns the URL path to the blog service oauth HTTP endpoint.
func OauthBlogPath() string {
	return "/oauth/redirect"
}

// JWTBlogPath returns the URL path to the blog service jwt HTTP endpoint.
func JWTBlogPath() string {
	return "/oauth"
}
