// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// RefreshAccessToken request
	RefreshAccessToken(ctx context.Context, params *RefreshAccessTokenParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Authenticate request with any body
	AuthenticateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTweetByID request
	GetTweetByID(ctx context.Context, tweetId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUsers request
	GetUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostUser request with any body
	PostUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostUser(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUserByID request
	GetUserByID(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PatchUser request with any body
	PatchUserWithBody(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PatchUser(ctx context.Context, userId int, body PatchUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) RefreshAccessToken(ctx context.Context, params *RefreshAccessTokenParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRefreshAccessTokenRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AuthenticateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAuthenticateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTweetByID(ctx context.Context, tweetId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTweetByIDRequest(c.Server, tweetId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUser(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUserByID(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserByIDRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchUserWithBody(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchUserRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PatchUser(ctx context.Context, userId int, body PatchUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPatchUserRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewRefreshAccessTokenRequest generates requests for RefreshAccessToken
func NewRefreshAccessTokenRequest(server string, params *RefreshAccessTokenParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/oauth2/refresh_token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "refresh_token", runtime.ParamLocationQuery, params.RefreshToken); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAuthenticateRequestWithBody generates requests for Authenticate with any type of body
func NewAuthenticateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/oauth2/token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetTweetByIDRequest generates requests for GetTweetByID
func NewGetTweetByIDRequest(server string, tweetId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "tweetId", runtime.ParamLocationPath, tweetId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tweets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetUsersRequest generates requests for GetUsers
func NewGetUsersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostUserRequest calls the generic PostUser builder with application/json body
func NewPostUserRequest(server string, body PostUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostUserRequestWithBody(server, "application/json", bodyReader)
}

// NewPostUserRequestWithBody generates requests for PostUser with any type of body
func NewPostUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetUserByIDRequest generates requests for GetUserByID
func NewGetUserByIDRequest(server string, userId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPatchUserRequest calls the generic PatchUser builder with application/json body
func NewPatchUserRequest(server string, userId int, body PatchUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPatchUserRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewPatchUserRequestWithBody generates requests for PatchUser with any type of body
func NewPatchUserRequestWithBody(server string, userId int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// RefreshAccessToken request
	RefreshAccessTokenWithResponse(ctx context.Context, params *RefreshAccessTokenParams, reqEditors ...RequestEditorFn) (*RefreshAccessTokenResponse, error)

	// Authenticate request with any body
	AuthenticateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error)

	// GetTweetByID request
	GetTweetByIDWithResponse(ctx context.Context, tweetId int, reqEditors ...RequestEditorFn) (*GetTweetByIDResponse, error)

	// GetUsers request
	GetUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsersResponse, error)

	// PostUser request with any body
	PostUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUserResponse, error)

	PostUserWithResponse(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUserResponse, error)

	// GetUserByID request
	GetUserByIDWithResponse(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*GetUserByIDResponse, error)

	// PatchUser request with any body
	PatchUserWithBodyWithResponse(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchUserResponse, error)

	PatchUserWithResponse(ctx context.Context, userId int, body PatchUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchUserResponse, error)
}

type RefreshAccessTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AuthenticationResponse
}

// Status returns HTTPResponse.Status
func (r RefreshAccessTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RefreshAccessTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AuthenticateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AuthenticationResponse
}

// Status returns HTTPResponse.Status
func (r AuthenticateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AuthenticateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTweetByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *map[string]interface{}
	JSON500      *ApiError
}

// Status returns HTTPResponse.Status
func (r GetTweetByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTweetByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]User
}

// Status returns HTTPResponse.Status
func (r GetUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
}

// Status returns HTTPResponse.Status
func (r GetUserByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PatchUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PatchUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PatchUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// RefreshAccessTokenWithResponse request returning *RefreshAccessTokenResponse
func (c *ClientWithResponses) RefreshAccessTokenWithResponse(ctx context.Context, params *RefreshAccessTokenParams, reqEditors ...RequestEditorFn) (*RefreshAccessTokenResponse, error) {
	rsp, err := c.RefreshAccessToken(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRefreshAccessTokenResponse(rsp)
}

// AuthenticateWithBodyWithResponse request with arbitrary body returning *AuthenticateResponse
func (c *ClientWithResponses) AuthenticateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error) {
	rsp, err := c.AuthenticateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAuthenticateResponse(rsp)
}

// GetTweetByIDWithResponse request returning *GetTweetByIDResponse
func (c *ClientWithResponses) GetTweetByIDWithResponse(ctx context.Context, tweetId int, reqEditors ...RequestEditorFn) (*GetTweetByIDResponse, error) {
	rsp, err := c.GetTweetByID(ctx, tweetId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTweetByIDResponse(rsp)
}

// GetUsersWithResponse request returning *GetUsersResponse
func (c *ClientWithResponses) GetUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsersResponse, error) {
	rsp, err := c.GetUsers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersResponse(rsp)
}

// PostUserWithBodyWithResponse request with arbitrary body returning *PostUserResponse
func (c *ClientWithResponses) PostUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUserResponse, error) {
	rsp, err := c.PostUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUserResponse(rsp)
}

func (c *ClientWithResponses) PostUserWithResponse(ctx context.Context, body PostUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUserResponse, error) {
	rsp, err := c.PostUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUserResponse(rsp)
}

// GetUserByIDWithResponse request returning *GetUserByIDResponse
func (c *ClientWithResponses) GetUserByIDWithResponse(ctx context.Context, userId int, reqEditors ...RequestEditorFn) (*GetUserByIDResponse, error) {
	rsp, err := c.GetUserByID(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserByIDResponse(rsp)
}

// PatchUserWithBodyWithResponse request with arbitrary body returning *PatchUserResponse
func (c *ClientWithResponses) PatchUserWithBodyWithResponse(ctx context.Context, userId int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PatchUserResponse, error) {
	rsp, err := c.PatchUserWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchUserResponse(rsp)
}

func (c *ClientWithResponses) PatchUserWithResponse(ctx context.Context, userId int, body PatchUserJSONRequestBody, reqEditors ...RequestEditorFn) (*PatchUserResponse, error) {
	rsp, err := c.PatchUser(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePatchUserResponse(rsp)
}

// ParseRefreshAccessTokenResponse parses an HTTP response from a RefreshAccessTokenWithResponse call
func ParseRefreshAccessTokenResponse(rsp *http.Response) (*RefreshAccessTokenResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RefreshAccessTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AuthenticationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseAuthenticateResponse parses an HTTP response from a AuthenticateWithResponse call
func ParseAuthenticateResponse(rsp *http.Response) (*AuthenticateResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AuthenticateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AuthenticationResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetTweetByIDResponse parses an HTTP response from a GetTweetByIDWithResponse call
func ParseGetTweetByIDResponse(rsp *http.Response) (*GetTweetByIDResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTweetByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ApiError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetUsersResponse parses an HTTP response from a GetUsersWithResponse call
func ParseGetUsersResponse(rsp *http.Response) (*GetUsersResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostUserResponse parses an HTTP response from a PostUserWithResponse call
func ParsePostUserResponse(rsp *http.Response) (*PostUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetUserByIDResponse parses an HTTP response from a GetUserByIDWithResponse call
func ParseGetUserByIDResponse(rsp *http.Response) (*GetUserByIDResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePatchUserResponse parses an HTTP response from a PatchUserWithResponse call
func ParsePatchUserResponse(rsp *http.Response) (*PatchUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PatchUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
