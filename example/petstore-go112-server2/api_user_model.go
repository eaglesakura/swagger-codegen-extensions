package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Create user
// This can only be done by the logged in user.
type UserApiCreateUserPostRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Created user object
	Body *User
}

// Getter for Endpoint
func (it *UserApiCreateUserPostRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiCreateUserPostRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiCreateUserPostRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiCreateUserPostRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiCreateUserPostRequest) Valid() error {

	if it.Body == nil {
		return xerrors.Errorf("required parameter(Body) read failed")
	}

	return nil
}

// Set Body from non pointer
func (it *UserApiCreateUserPostRequest) SetBody(newBody User) {
	it.Body = &newBody
}

// Remove Body property.
func (it *UserApiCreateUserPostRequest) RemoveBody() {
	it.Body = nil
}

// Require value of Body
func (it *UserApiCreateUserPostRequest) RequireBody() User {
	if it.Body == nil {
		panic(xerrors.Errorf("UserApi.Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *UserApiCreateUserPostRequest) GetBody() User {
	if it.Body != nil {
		return *it.Body
	}
	result := new(User)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUserPostRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	return body, body, resp, nil
}

// Build http request
func (it *UserApiCreateUserPostRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.Body != nil {
		body = bodyToReader(it.Body)
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUserPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiCreateUserPostRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Creates list of users with given input array
//
type UserApiCreateUsersWithArrayInputPostRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// List of user object
	Body *[]User
}

// Getter for Endpoint
func (it *UserApiCreateUsersWithArrayInputPostRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiCreateUsersWithArrayInputPostRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiCreateUsersWithArrayInputPostRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiCreateUsersWithArrayInputPostRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiCreateUsersWithArrayInputPostRequest) Valid() error {

	if it.Body == nil {
		return xerrors.Errorf("required parameter(Body) read failed")
	}

	return nil
}

// Set Body from non pointer
func (it *UserApiCreateUsersWithArrayInputPostRequest) SetBody(newBody []User) {
	it.Body = &newBody
}

// Remove Body property.
func (it *UserApiCreateUsersWithArrayInputPostRequest) RemoveBody() {
	it.Body = nil
}

// Require value of Body
func (it *UserApiCreateUsersWithArrayInputPostRequest) RequireBody() []User {
	if it.Body == nil {
		panic(xerrors.Errorf("UserApi.Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *UserApiCreateUsersWithArrayInputPostRequest) GetBody() []User {
	if it.Body != nil {
		return *it.Body
	}
	result := new([]User)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUsersWithArrayInputPostRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	return body, body, resp, nil
}

// Build http request
func (it *UserApiCreateUsersWithArrayInputPostRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user/createWithArray")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.Body != nil {
		body = bodyToReader(it.Body)
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUsersWithArrayInputPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiCreateUsersWithArrayInputPostRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Creates list of users with given input array
//
type UserApiCreateUsersWithListInputPostRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// List of user object
	Body *[]User
}

// Getter for Endpoint
func (it *UserApiCreateUsersWithListInputPostRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiCreateUsersWithListInputPostRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiCreateUsersWithListInputPostRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiCreateUsersWithListInputPostRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiCreateUsersWithListInputPostRequest) Valid() error {

	if it.Body == nil {
		return xerrors.Errorf("required parameter(Body) read failed")
	}

	return nil
}

// Set Body from non pointer
func (it *UserApiCreateUsersWithListInputPostRequest) SetBody(newBody []User) {
	it.Body = &newBody
}

// Remove Body property.
func (it *UserApiCreateUsersWithListInputPostRequest) RemoveBody() {
	it.Body = nil
}

// Require value of Body
func (it *UserApiCreateUsersWithListInputPostRequest) RequireBody() []User {
	if it.Body == nil {
		panic(xerrors.Errorf("UserApi.Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *UserApiCreateUsersWithListInputPostRequest) GetBody() []User {
	if it.Body != nil {
		return *it.Body
	}
	result := new([]User)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUsersWithListInputPostRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	return body, body, resp, nil
}

// Build http request
func (it *UserApiCreateUsersWithListInputPostRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user/createWithList")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.Body != nil {
		body = bodyToReader(it.Body)
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUsersWithListInputPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiCreateUsersWithListInputPostRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Delete user
// This can only be done by the logged in user.
type UserApiDeleteUserDeleteRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// The name that needs to be deleted
	Username *string
}

// Getter for Endpoint
func (it *UserApiDeleteUserDeleteRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiDeleteUserDeleteRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiDeleteUserDeleteRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiDeleteUserDeleteRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiDeleteUserDeleteRequest) Valid() error {

	if it.Username == nil {
		return xerrors.Errorf("required parameter(Username) read failed")
	}

	return nil
}

// Set Username from non pointer
func (it *UserApiDeleteUserDeleteRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
}

// Remove Username property.
func (it *UserApiDeleteUserDeleteRequest) RemoveUsername() {
	it.Username = nil
}

// Require value of Username
func (it *UserApiDeleteUserDeleteRequest) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf("UserApi.Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *UserApiDeleteUserDeleteRequest) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *UserApiDeleteUserDeleteRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	return body, body, resp, nil
}

// Build http request
func (it *UserApiDeleteUserDeleteRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user/{username}")
	method := strings.ToUpper("Delete")
	var body io.Reader

	if it.Username != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"username"+"}", url.PathEscape(fmt.Sprintf("%v", *it.Username)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiDeleteUserDeleteRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiDeleteUserDeleteRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Get user by user name
//
type UserApiGetUserByNameGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// The name that needs to be fetched. Use user1 for testing.
	Username *string
}

// Getter for Endpoint
func (it *UserApiGetUserByNameGetRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiGetUserByNameGetRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiGetUserByNameGetRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiGetUserByNameGetRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiGetUserByNameGetRequest) Valid() error {

	if it.Username == nil {
		return xerrors.Errorf("required parameter(Username) read failed")
	}

	return nil
}

// Set Username from non pointer
func (it *UserApiGetUserByNameGetRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
}

// Remove Username property.
func (it *UserApiGetUserByNameGetRequest) RemoveUsername() {
	it.Username = nil
}

// Require value of Username
func (it *UserApiGetUserByNameGetRequest) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf("UserApi.Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *UserApiGetUserByNameGetRequest) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *UserApiGetUserByNameGetRequest) Execute(ctx context.Context) (success *User, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	var model User
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed: %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *UserApiGetUserByNameGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user/{username}")
	method := strings.ToUpper("Get")
	var body io.Reader

	if it.Username != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"username"+"}", url.PathEscape(fmt.Sprintf("%v", *it.Username)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiGetUserByNameGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiGetUserByNameGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Logs user into the system
//
type UserApiLoginUserGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// The user name for login
	Username *string
	// The password for login in clear text
	Password *string
}

// Getter for Endpoint
func (it *UserApiLoginUserGetRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiLoginUserGetRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiLoginUserGetRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiLoginUserGetRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiLoginUserGetRequest) Valid() error {

	if it.Username == nil {
		return xerrors.Errorf("required parameter(Username) read failed")
	}

	if it.Password == nil {
		return xerrors.Errorf("required parameter(Password) read failed")
	}

	return nil
}

// Set Username from non pointer
func (it *UserApiLoginUserGetRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
}

// Remove Username property.
func (it *UserApiLoginUserGetRequest) RemoveUsername() {
	it.Username = nil
}

// Require value of Username
func (it *UserApiLoginUserGetRequest) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf("UserApi.Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *UserApiLoginUserGetRequest) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

// Set Password from non pointer
func (it *UserApiLoginUserGetRequest) SetPassword(newPassword string) {
	it.Password = &newPassword
}

// Remove Password property.
func (it *UserApiLoginUserGetRequest) RemovePassword() {
	it.Password = nil
}

// Require value of Password
func (it *UserApiLoginUserGetRequest) RequirePassword() string {
	if it.Password == nil {
		panic(xerrors.Errorf("UserApi.Password is nil"))
	}
	return *it.Password
}

// Get value of Password / or default
func (it *UserApiLoginUserGetRequest) GetPassword() string {
	if it.Password != nil {
		return *it.Password
	}
	result := new(string)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *UserApiLoginUserGetRequest) Execute(ctx context.Context) (success *string, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	var model string
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed: %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *UserApiLoginUserGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user/login")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	if it.Username != nil {
		query.Set("username", url.QueryEscape(fmt.Sprintf("%v", *it.Username)))
	}
	if it.Password != nil {
		query.Set("password", url.QueryEscape(fmt.Sprintf("%v", *it.Password)))
	}
	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiLoginUserGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiLoginUserGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Logs out current logged in user session
//
type UserApiLogoutUserGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

// Getter for Endpoint
func (it *UserApiLogoutUserGetRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiLogoutUserGetRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiLogoutUserGetRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiLogoutUserGetRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiLogoutUserGetRequest) Valid() error {

	return nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiLogoutUserGetRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	return body, body, resp, nil
}

// Build http request
func (it *UserApiLogoutUserGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user/logout")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiLogoutUserGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiLogoutUserGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Updated user
// This can only be done by the logged in user.
type UserApiUpdateUserPutRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// name that need to be updated
	Username *string
	// Updated user object
	Body *User
}

// Getter for Endpoint
func (it *UserApiUpdateUserPutRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *UserApiUpdateUserPutRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *UserApiUpdateUserPutRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *UserApiUpdateUserPutRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *UserApiUpdateUserPutRequest) Valid() error {

	if it.Username == nil {
		return xerrors.Errorf("required parameter(Username) read failed")
	}

	if it.Body == nil {
		return xerrors.Errorf("required parameter(Body) read failed")
	}

	return nil
}

// Set Username from non pointer
func (it *UserApiUpdateUserPutRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
}

// Remove Username property.
func (it *UserApiUpdateUserPutRequest) RemoveUsername() {
	it.Username = nil
}

// Require value of Username
func (it *UserApiUpdateUserPutRequest) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf("UserApi.Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *UserApiUpdateUserPutRequest) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

// Set Body from non pointer
func (it *UserApiUpdateUserPutRequest) SetBody(newBody User) {
	it.Body = &newBody
}

// Remove Body property.
func (it *UserApiUpdateUserPutRequest) RemoveBody() {
	it.Body = nil
}

// Require value of Body
func (it *UserApiUpdateUserPutRequest) RequireBody() User {
	if it.Body == nil {
		panic(xerrors.Errorf("UserApi.Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *UserApiUpdateUserPutRequest) GetBody() User {
	if it.Body != nil {
		return *it.Body
	}
	result := new(User)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *UserApiUpdateUserPutRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
	resp, err = it.Fetch(ctx)
	if resp != nil && resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http fetch failed, : %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response: %w", err)
	}

	return body, body, resp, nil
}

// Build http request
func (it *UserApiUpdateUserPutRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/user/{username}")
	method := strings.ToUpper("Put")
	var body io.Reader

	if it.Username != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"username"+"}", url.PathEscape(fmt.Sprintf("%v", *it.Username)))
	}
	if it.Body != nil {
		body = bodyToReader(it.Body)
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiUpdateUserPutRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed: %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed: %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *UserApiUpdateUserPutRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}
