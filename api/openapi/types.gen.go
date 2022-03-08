// Package Openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package Openapi

// Tweet defines model for Tweet.
type Tweet struct {
	Id     int64   `json:"id"`
	Text   *string `json:"text,omitempty"`
	UserId *int64  `json:"user_id,omitempty"`
}

// User defines model for User.
type User struct {
	Id   int64   `json:"id"`
	Name *string `json:"name,omitempty"`
}

