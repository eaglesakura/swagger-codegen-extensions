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

type UserApi struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

func NewUserApi() *UserApi {
	return &UserApi{
		Endpoint: "",
	}
}

// Create user
// This can only be done by the logged in user.
type UserApiCreateUserPostRequest struct {
	api *UserApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Created user object
	Body *User
}

// Set Body from non pointer
func (it *UserApiCreateUserPostRequest) SetBody(newBody User) {
	it.Body = &newBody
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	return body, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUserPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.Body != nil {
		body = bodyToReader(it.Body)
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
// body: Created user object
func (it *UserApi) CreateUserPost(builder func(*UserApiCreateUserPostRequest)) *UserApiCreateUserPostRequest {
	result := &UserApiCreateUserPostRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Creates list of users with given input array
//
type UserApiCreateUsersWithArrayInputPostRequest struct {
	api *UserApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// List of user object
	Body *[]User
}

// Set Body from non pointer
func (it *UserApiCreateUsersWithArrayInputPostRequest) SetBody(newBody []User) {
	it.Body = &newBody
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	return body, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUsersWithArrayInputPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user/createWithArray")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.Body != nil {
		body = bodyToReader(it.Body)
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
// body: List of user object
func (it *UserApi) CreateUsersWithArrayInputPost(builder func(*UserApiCreateUsersWithArrayInputPostRequest)) *UserApiCreateUsersWithArrayInputPostRequest {
	result := &UserApiCreateUsersWithArrayInputPostRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Creates list of users with given input array
//
type UserApiCreateUsersWithListInputPostRequest struct {
	api *UserApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// List of user object
	Body *[]User
}

// Set Body from non pointer
func (it *UserApiCreateUsersWithListInputPostRequest) SetBody(newBody []User) {
	it.Body = &newBody
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	return body, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiCreateUsersWithListInputPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user/createWithList")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.Body != nil {
		body = bodyToReader(it.Body)
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
// body: List of user object
func (it *UserApi) CreateUsersWithListInputPost(builder func(*UserApiCreateUsersWithListInputPostRequest)) *UserApiCreateUsersWithListInputPostRequest {
	result := &UserApiCreateUsersWithListInputPostRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Delete user
// This can only be done by the logged in user.
type UserApiDeleteUserDeleteRequest struct {
	api *UserApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// The name that needs to be deleted
	Username *string
}

// Set Username from non pointer
func (it *UserApiDeleteUserDeleteRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	return body, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiDeleteUserDeleteRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user/{username}")
	method := strings.ToUpper("Delete")
	var body io.Reader

	if it.Username != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"username"+"}", url.PathEscape(fmt.Sprintf("%v", *it.Username)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
// username: The name that needs to be deleted
func (it *UserApi) DeleteUserDelete(builder func(*UserApiDeleteUserDeleteRequest)) *UserApiDeleteUserDeleteRequest {
	result := &UserApiDeleteUserDeleteRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Get user by user name
//
type UserApiGetUserByNameGetRequest struct {
	api *UserApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// The name that needs to be fetched. Use user1 for testing.
	Username *string
}

// Set Username from non pointer
func (it *UserApiGetUserByNameGetRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	var model User
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiGetUserByNameGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user/{username}")
	method := strings.ToUpper("Get")
	var body io.Reader

	if it.Username != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"username"+"}", url.PathEscape(fmt.Sprintf("%v", *it.Username)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
// username: The name that needs to be fetched. Use user1 for testing.
func (it *UserApi) GetUserByNameGet(builder func(*UserApiGetUserByNameGetRequest)) *UserApiGetUserByNameGetRequest {
	result := &UserApiGetUserByNameGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Logs user into the system
//
type UserApiLoginUserGetRequest struct {
	api *UserApi

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

// Set Username from non pointer
func (it *UserApiLoginUserGetRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
}

// Set Password from non pointer
func (it *UserApiLoginUserGetRequest) SetPassword(newPassword string) {
	it.Password = &newPassword
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	var model string
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiLoginUserGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user/login")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	if it.Username != nil {
		query.Set("username", url.QueryEscape(fmt.Sprintf("%v", *it.Username)))
	}
	if it.Password != nil {
		query.Set("password", url.QueryEscape(fmt.Sprintf("%v", *it.Password)))
	}

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
// username: The user name for login
// password: The password for login in clear text
func (it *UserApi) LoginUserGet(builder func(*UserApiLoginUserGetRequest)) *UserApiLoginUserGetRequest {
	result := &UserApiLoginUserGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Logs out current logged in user session
//
type UserApiLogoutUserGetRequest struct {
	api *UserApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	return body, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiLogoutUserGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user/logout")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
func (it *UserApi) LogoutUserGet(builder func(*UserApiLogoutUserGetRequest)) *UserApiLogoutUserGetRequest {
	result := &UserApiLogoutUserGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Updated user
// This can only be done by the logged in user.
type UserApiUpdateUserPutRequest struct {
	api *UserApi

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

// Set Username from non pointer
func (it *UserApiUpdateUserPutRequest) SetUsername(newUsername string) {
	it.Username = &newUsername
}

// Set Body from non pointer
func (it *UserApiUpdateUserPutRequest) SetBody(newBody User) {
	it.Body = &newBody
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	return body, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *UserApiUpdateUserPutRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/user/{username}")
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
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()

	if it.Intercept != nil {
		it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

// New request with Parameters.
// username: name that need to be updated
// body: Updated user object
func (it *UserApi) UpdateUserPut(builder func(*UserApiUpdateUserPutRequest)) *UserApiUpdateUserPutRequest {
	result := &UserApiUpdateUserPutRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

func (it *UserApi) this_is_call_dummy() {
	url.Parse("")
	xerrors.Errorf("")
	strings.ToUpper("")
	fmt.Sprintf("%v", "")
	io.ReadAtLeast(nil, nil, 0)
	json.Unmarshal(nil, nil)
}
