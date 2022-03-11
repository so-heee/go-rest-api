// Package Openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package Openapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Authenticate with the admin server.
	// (POST /auth)
	Authenticate(ctx echo.Context) error
	// Get tweet by ID.
	// (GET /tweets/{tweetId})
	GetTweetByID(ctx echo.Context, tweetId int) error
	// Get user by ID.
	// (GET /users/{userId})
	GetUserByID(ctx echo.Context, userId int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Authenticate converts echo context to params.
func (w *ServerInterfaceWrapper) Authenticate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Authenticate(ctx)
	return err
}

// GetTweetByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetTweetByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tweetId" -------------
	var tweetId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "tweetId", runtime.ParamLocationPath, ctx.Param("tweetId"), &tweetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tweetId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTweetByID(ctx, tweetId)
	return err
}

// GetUserByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUserByID(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/auth", wrapper.Authenticate)
	router.GET(baseURL+"/tweets/:tweetId", wrapper.GetTweetByID)
	router.GET(baseURL+"/users/:userId", wrapper.GetUserByID)

}

