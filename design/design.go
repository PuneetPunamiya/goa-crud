package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("blog", func() {
	Title("Blog Service")
	Description("Service to perform CRUD operations using goa")
	Server("blog", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// APIKeyAuth defines a security scheme that uses API keys.
var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("Secures endpoint by requiring an API key.")
})

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during signin")
	Scope("api:read", "Read-only access")
})

// OAuth2Auth defines a security scheme that uses OAuth2 tokens.
var OAuth2Auth = OAuth2Security("oauth2", func() {
	AuthorizationCodeFlow("http://goa.design/authorization", "http://goa.design/token", "http://goa.design/refresh")
	Description(`Secures endpoint by requiring a valid OAuth2 token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

var _ = Service("blog", func() {
	Description("The blog service gives blog details.")

	cors.Origin("/.*localhost/", func() {
		cors.Headers("Content-Type")
		cors.Methods("GET", "POST")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})

	//Method to post new blog
	Method("create", func() {
		Description("Add new blog and return its ID.")
		Payload(Blog)
		Result(Blog)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	//Method to get all existing blogs
	Method("list", func() {
		Description("List all entries")
		Result(ArrayOf(StoredBlogs))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	//Method to remove a particular blog
	Method("remove", func() {
		Description("Remove blog from storage")
		Payload(func() {
			Field(1, "id", UInt32, "ID of blog to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Blog not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
		})
	})

	//Method to update a specific blog
	Method("update", func() {
		Description("Updating the existing blog")
		Payload(func() {
			Field(1, "id", UInt32, "ID of blog to be updated")
			Field(2, "name", String, "Details of blog to be updated")
			Field(3, "comments", ArrayOf(comments), "Comments to be updated")
			Required("name", "comments")
		})
		HTTP(func() {
			PATCH("/{id}")
			Response(StatusNoContent)
		})

	})

	//Method to add a new comment to an existing blog
	Method("add", func() {
		Description("Add new blog and return its ID.")
		Payload(new_comment)
		Result(new_comment)
		HTTP(func() {
			POST("/{id}/comments/")
			Response(StatusCreated)
		})
	})

	//Method to get a particular blog based on id
	Method("show", func() {
		Description("Show blog based on the id given")
		Payload(Blog)
		Result(Blog)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("oauth", func() {
		Description("Github authentication to post a new blog")
		Payload(func() {
			Field(1, "token", String, "Access github token")
		})
		Result(String)
		HTTP(func() {
			POST("/oauth/redirect")
			Response(StatusCreated)
		})

	})

	Method("jwt", func() {
		Description("Getting auth")
		Payload(func() {
			Field(1, "auth", String, "Access github token")
		})
		Result(String)
		HTTP(func() {
			POST("/oauth")
			Header("auth:X-Authorization") // JWT token passed in "X-Authorization" header

			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
