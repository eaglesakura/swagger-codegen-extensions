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

// Add a new pet to the store
//
type PetApiAddPetPostRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Pet object that needs to be added to the store
	Body *Pet
}

// Getter for Endpoint
func (it *PetApiAddPetPostRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiAddPetPostRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiAddPetPostRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiAddPetPostRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiAddPetPostRequest) Valid() error {

	if it.Body == nil {
		return xerrors.Errorf("required parameter(Body) read failed")
	}

	return nil
}

// Set Body from non pointer
func (it *PetApiAddPetPostRequest) SetBody(newBody Pet) {
	it.Body = &newBody
}

// Remove Body property.
func (it *PetApiAddPetPostRequest) RemoveBody() {
	it.Body = nil
}

// Require value of Body
func (it *PetApiAddPetPostRequest) RequireBody() Pet {
	if it.Body == nil {
		panic(xerrors.Errorf("PetApi.Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *PetApiAddPetPostRequest) GetBody() Pet {
	if it.Body != nil {
		return *it.Body
	}
	result := new(Pet)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiAddPetPostRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
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
func (it *PetApiAddPetPostRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet")
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
func (it *PetApiAddPetPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiAddPetPostRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Deletes a pet
//
type PetApiDeletePetDeleteRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Pet id to delete
	PetId *int64
	//
	ApiKey *string
}

// Getter for Endpoint
func (it *PetApiDeletePetDeleteRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiDeletePetDeleteRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiDeletePetDeleteRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiDeletePetDeleteRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiDeletePetDeleteRequest) Valid() error {

	if it.PetId == nil {
		return xerrors.Errorf("required parameter(PetId) read failed")
	}

	return nil
}

// Set PetId from non pointer
func (it *PetApiDeletePetDeleteRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Remove PetId property.
func (it *PetApiDeletePetDeleteRequest) RemovePetId() {
	it.PetId = nil
}

// Require value of PetId
func (it *PetApiDeletePetDeleteRequest) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf("PetApi.PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *PetApiDeletePetDeleteRequest) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Set ApiKey from non pointer
func (it *PetApiDeletePetDeleteRequest) SetApiKey(newApiKey string) {
	it.ApiKey = &newApiKey
}

// Remove ApiKey property.
func (it *PetApiDeletePetDeleteRequest) RemoveApiKey() {
	it.ApiKey = nil
}

// Require value of ApiKey
func (it *PetApiDeletePetDeleteRequest) RequireApiKey() string {
	if it.ApiKey == nil {
		panic(xerrors.Errorf("PetApi.ApiKey is nil"))
	}
	return *it.ApiKey
}

// Get value of ApiKey / or default
func (it *PetApiDeletePetDeleteRequest) GetApiKey() string {
	if it.ApiKey != nil {
		return *it.ApiKey
	}
	result := new(string)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiDeletePetDeleteRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
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
func (it *PetApiDeletePetDeleteRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet/{petId}")
	method := strings.ToUpper("Delete")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	if it.ApiKey != nil {
		request.Header.Set("api_key", fmt.Sprintf("%v", *it.ApiKey))
	}
	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *PetApiDeletePetDeleteRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiDeletePetDeleteRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Finds Pets by status
// Multiple status values can be provided with comma separated strings
type PetApiFindPetsByStatusGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Status values that need to be considered for filter
	Status *[]string
}

// Getter for Endpoint
func (it *PetApiFindPetsByStatusGetRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiFindPetsByStatusGetRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiFindPetsByStatusGetRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiFindPetsByStatusGetRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiFindPetsByStatusGetRequest) Valid() error {

	if it.Status == nil {
		return xerrors.Errorf("required parameter(Status) read failed")
	}

	return nil
}

// Set Status from non pointer
func (it *PetApiFindPetsByStatusGetRequest) SetStatus(newStatus []string) {
	it.Status = &newStatus
}

// Remove Status property.
func (it *PetApiFindPetsByStatusGetRequest) RemoveStatus() {
	it.Status = nil
}

// Require value of Status
func (it *PetApiFindPetsByStatusGetRequest) RequireStatus() []string {
	if it.Status == nil {
		panic(xerrors.Errorf("PetApi.Status is nil"))
	}
	return *it.Status
}

// Get value of Status / or default
func (it *PetApiFindPetsByStatusGetRequest) GetStatus() []string {
	if it.Status != nil {
		return *it.Status
	}
	result := new([]string)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiFindPetsByStatusGetRequest) Execute(ctx context.Context) (success *[]Pet, responseBody []byte, resp *http.Response, err error) {
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

	var model []Pet
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed: %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *PetApiFindPetsByStatusGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet/findByStatus")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	if it.Status != nil {
		query.Set("status", url.QueryEscape(fmt.Sprintf("%v", *it.Status)))
	}
	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *PetApiFindPetsByStatusGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiFindPetsByStatusGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Finds Pets by tags
// Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
type PetApiFindPetsByTagsGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Tags to filter by
	Tags *[]string
}

// Getter for Endpoint
func (it *PetApiFindPetsByTagsGetRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiFindPetsByTagsGetRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiFindPetsByTagsGetRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiFindPetsByTagsGetRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiFindPetsByTagsGetRequest) Valid() error {

	if it.Tags == nil {
		return xerrors.Errorf("required parameter(Tags) read failed")
	}

	return nil
}

// Set Tags from non pointer
func (it *PetApiFindPetsByTagsGetRequest) SetTags(newTags []string) {
	it.Tags = &newTags
}

// Remove Tags property.
func (it *PetApiFindPetsByTagsGetRequest) RemoveTags() {
	it.Tags = nil
}

// Require value of Tags
func (it *PetApiFindPetsByTagsGetRequest) RequireTags() []string {
	if it.Tags == nil {
		panic(xerrors.Errorf("PetApi.Tags is nil"))
	}
	return *it.Tags
}

// Get value of Tags / or default
func (it *PetApiFindPetsByTagsGetRequest) GetTags() []string {
	if it.Tags != nil {
		return *it.Tags
	}
	result := new([]string)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiFindPetsByTagsGetRequest) Execute(ctx context.Context) (success *[]Pet, responseBody []byte, resp *http.Response, err error) {
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

	var model []Pet
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed: %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *PetApiFindPetsByTagsGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet/findByTags")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed: %w", err)
	}

	query := request.URL.Query()

	if it.Tags != nil {
		query.Set("tags", url.QueryEscape(fmt.Sprintf("%v", *it.Tags)))
	}
	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *PetApiFindPetsByTagsGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiFindPetsByTagsGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Find pet by ID
// Returns a single pet
type PetApiGetPetByIdGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of pet to return
	PetId *int64
}

// Getter for Endpoint
func (it *PetApiGetPetByIdGetRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiGetPetByIdGetRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiGetPetByIdGetRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiGetPetByIdGetRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiGetPetByIdGetRequest) Valid() error {

	if it.PetId == nil {
		return xerrors.Errorf("required parameter(PetId) read failed")
	}

	return nil
}

// Set PetId from non pointer
func (it *PetApiGetPetByIdGetRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Remove PetId property.
func (it *PetApiGetPetByIdGetRequest) RemovePetId() {
	it.PetId = nil
}

// Require value of PetId
func (it *PetApiGetPetByIdGetRequest) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf("PetApi.PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *PetApiGetPetByIdGetRequest) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiGetPetByIdGetRequest) Execute(ctx context.Context) (success *Pet, responseBody []byte, resp *http.Response, err error) {
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

	var model Pet
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed: %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *PetApiGetPetByIdGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet/{petId}")
	method := strings.ToUpper("Get")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
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
func (it *PetApiGetPetByIdGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiGetPetByIdGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Update an existing pet
//
type PetApiUpdatePetPutRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Pet object that needs to be added to the store
	Body *Pet
}

// Getter for Endpoint
func (it *PetApiUpdatePetPutRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiUpdatePetPutRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiUpdatePetPutRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiUpdatePetPutRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiUpdatePetPutRequest) Valid() error {

	if it.Body == nil {
		return xerrors.Errorf("required parameter(Body) read failed")
	}

	return nil
}

// Set Body from non pointer
func (it *PetApiUpdatePetPutRequest) SetBody(newBody Pet) {
	it.Body = &newBody
}

// Remove Body property.
func (it *PetApiUpdatePetPutRequest) RemoveBody() {
	it.Body = nil
}

// Require value of Body
func (it *PetApiUpdatePetPutRequest) RequireBody() Pet {
	if it.Body == nil {
		panic(xerrors.Errorf("PetApi.Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *PetApiUpdatePetPutRequest) GetBody() Pet {
	if it.Body != nil {
		return *it.Body
	}
	result := new(Pet)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiUpdatePetPutRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
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
func (it *PetApiUpdatePetPutRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet")
	method := strings.ToUpper("Put")
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
func (it *PetApiUpdatePetPutRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiUpdatePetPutRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Updates a pet in the store with form data
//
type PetApiUpdatePetWithFormPostRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of pet that needs to be updated
	PetId *int64
	// Updated name of the pet
	Name *string
	// Updated status of the pet
	Status *string
}

// Getter for Endpoint
func (it *PetApiUpdatePetWithFormPostRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiUpdatePetWithFormPostRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiUpdatePetWithFormPostRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiUpdatePetWithFormPostRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiUpdatePetWithFormPostRequest) Valid() error {

	if it.PetId == nil {
		return xerrors.Errorf("required parameter(PetId) read failed")
	}

	return nil
}

// Set PetId from non pointer
func (it *PetApiUpdatePetWithFormPostRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Remove PetId property.
func (it *PetApiUpdatePetWithFormPostRequest) RemovePetId() {
	it.PetId = nil
}

// Require value of PetId
func (it *PetApiUpdatePetWithFormPostRequest) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf("PetApi.PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *PetApiUpdatePetWithFormPostRequest) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Set Name from non pointer
func (it *PetApiUpdatePetWithFormPostRequest) SetName(newName string) {
	it.Name = &newName
}

// Remove Name property.
func (it *PetApiUpdatePetWithFormPostRequest) RemoveName() {
	it.Name = nil
}

// Require value of Name
func (it *PetApiUpdatePetWithFormPostRequest) RequireName() string {
	if it.Name == nil {
		panic(xerrors.Errorf("PetApi.Name is nil"))
	}
	return *it.Name
}

// Get value of Name / or default
func (it *PetApiUpdatePetWithFormPostRequest) GetName() string {
	if it.Name != nil {
		return *it.Name
	}
	result := new(string)
	return *result
}

// Set Status from non pointer
func (it *PetApiUpdatePetWithFormPostRequest) SetStatus(newStatus string) {
	it.Status = &newStatus
}

// Remove Status property.
func (it *PetApiUpdatePetWithFormPostRequest) RemoveStatus() {
	it.Status = nil
}

// Require value of Status
func (it *PetApiUpdatePetWithFormPostRequest) RequireStatus() string {
	if it.Status == nil {
		panic(xerrors.Errorf("PetApi.Status is nil"))
	}
	return *it.Status
}

// Get value of Status / or default
func (it *PetApiUpdatePetWithFormPostRequest) GetStatus() string {
	if it.Status != nil {
		return *it.Status
	}
	result := new(string)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiUpdatePetWithFormPostRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
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
func (it *PetApiUpdatePetWithFormPostRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet/{petId}")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
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
func (it *PetApiUpdatePetWithFormPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiUpdatePetWithFormPostRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// uploads an image
//
type PetApiUploadFilePostRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of pet to update
	PetId *int64
	// Additional data to pass to server
	AdditionalMetadata *string
	// file to upload
	File *io.Reader
}

// Getter for Endpoint
func (it *PetApiUploadFilePostRequest) GetEndpoint() string {
	return it.Endpoint
}

// Setter for Endpoint
func (it *PetApiUploadFilePostRequest) SetEndpoint(newEndpoint string) {
	it.Endpoint = newEndpoint
}

// Getter for Intercept function
func (it *PetApiUploadFilePostRequest) GetIntercept() func(client *http.Client, request *http.Request) (*http.Client, *http.Request) {
	return it.Intercept
}

// Setter for Intercept function
func (it *PetApiUploadFilePostRequest) SetIntercept(newIntercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)) {
	it.Intercept = newIntercept
}

// Validation parameter
func (it *PetApiUploadFilePostRequest) Valid() error {

	if it.PetId == nil {
		return xerrors.Errorf("required parameter(PetId) read failed")
	}

	return nil
}

// Set PetId from non pointer
func (it *PetApiUploadFilePostRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Remove PetId property.
func (it *PetApiUploadFilePostRequest) RemovePetId() {
	it.PetId = nil
}

// Require value of PetId
func (it *PetApiUploadFilePostRequest) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf("PetApi.PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *PetApiUploadFilePostRequest) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Set AdditionalMetadata from non pointer
func (it *PetApiUploadFilePostRequest) SetAdditionalMetadata(newAdditionalMetadata string) {
	it.AdditionalMetadata = &newAdditionalMetadata
}

// Remove AdditionalMetadata property.
func (it *PetApiUploadFilePostRequest) RemoveAdditionalMetadata() {
	it.AdditionalMetadata = nil
}

// Require value of AdditionalMetadata
func (it *PetApiUploadFilePostRequest) RequireAdditionalMetadata() string {
	if it.AdditionalMetadata == nil {
		panic(xerrors.Errorf("PetApi.AdditionalMetadata is nil"))
	}
	return *it.AdditionalMetadata
}

// Get value of AdditionalMetadata / or default
func (it *PetApiUploadFilePostRequest) GetAdditionalMetadata() string {
	if it.AdditionalMetadata != nil {
		return *it.AdditionalMetadata
	}
	result := new(string)
	return *result
}

// Set File from non pointer
func (it *PetApiUploadFilePostRequest) SetFile(newFile io.Reader) {
	it.File = &newFile
}

// Remove File property.
func (it *PetApiUploadFilePostRequest) RemoveFile() {
	it.File = nil
}

// Require value of File
func (it *PetApiUploadFilePostRequest) RequireFile() io.Reader {
	if it.File == nil {
		panic(xerrors.Errorf("PetApi.File is nil"))
	}
	return *it.File
}

// Get value of File / or default
func (it *PetApiUploadFilePostRequest) GetFile() io.Reader {
	if it.File != nil {
		return *it.File
	}
	result := new(io.Reader)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *PetApiUploadFilePostRequest) Execute(ctx context.Context) (success *ApiResponse, responseBody []byte, resp *http.Response, err error) {
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

	var model ApiResponse
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed: %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *PetApiUploadFilePostRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/pet/{petId}/uploadImage")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
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
func (it *PetApiUploadFilePostRequest) Fetch(ctx context.Context) (*http.Response, error) {
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
func (it *PetApiUploadFilePostRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}
