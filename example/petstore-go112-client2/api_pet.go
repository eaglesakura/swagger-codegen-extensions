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

type PetApi struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

func NewPetApi() *PetApi {
	return &PetApi{
		Endpoint: "",
	}
}

// Add a new pet to the store
//
type PetApiAddPetPostRequest struct {
	api *PetApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Pet object that needs to be added to the store
	Body *Pet
}

// Set Body from non pointer
func (it *PetApiAddPetPostRequest) SetBody(newBody Pet) {
	it.Body = &newBody
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
func (it *PetApiAddPetPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet")
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
// body: Pet object that needs to be added to the store
func (it *PetApi) AddPetPost(builder func(*PetApiAddPetPostRequest)) *PetApiAddPetPostRequest {
	result := &PetApiAddPetPostRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Deletes a pet
//
type PetApiDeletePetDeleteRequest struct {
	api *PetApi

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

// Set PetId from non pointer
func (it *PetApiDeletePetDeleteRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Set ApiKey from non pointer
func (it *PetApiDeletePetDeleteRequest) SetApiKey(newApiKey string) {
	it.ApiKey = &newApiKey
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
func (it *PetApiDeletePetDeleteRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet/{petId}")
	method := strings.ToUpper("Delete")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	if it.ApiKey != nil {
		request.Header.Set("api_key", fmt.Sprintf("%v", *it.ApiKey))
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
// petId: Pet id to delete
// api_key:
func (it *PetApi) DeletePetDelete(builder func(*PetApiDeletePetDeleteRequest)) *PetApiDeletePetDeleteRequest {
	result := &PetApiDeletePetDeleteRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Finds Pets by status
// Multiple status values can be provided with comma separated strings
type PetApiFindPetsByStatusGetRequest struct {
	api *PetApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Status values that need to be considered for filter
	Status *[]string
}

// Set Status from non pointer
func (it *PetApiFindPetsByStatusGetRequest) SetStatus(newStatus []string) {
	it.Status = &newStatus
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	var model []Pet
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *PetApiFindPetsByStatusGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet/findByStatus")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	if it.Status != nil {
		query.Set("status", url.QueryEscape(fmt.Sprintf("%v", *it.Status)))
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
// status: Status values that need to be considered for filter
func (it *PetApi) FindPetsByStatusGet(builder func(*PetApiFindPetsByStatusGetRequest)) *PetApiFindPetsByStatusGetRequest {
	result := &PetApiFindPetsByStatusGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Finds Pets by tags
// Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
type PetApiFindPetsByTagsGetRequest struct {
	api *PetApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Tags to filter by
	Tags *[]string
}

// Set Tags from non pointer
func (it *PetApiFindPetsByTagsGetRequest) SetTags(newTags []string) {
	it.Tags = &newTags
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	var model []Pet
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *PetApiFindPetsByTagsGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet/findByTags")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	if it.Tags != nil {
		query.Set("tags", url.QueryEscape(fmt.Sprintf("%v", *it.Tags)))
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
// tags: Tags to filter by
func (it *PetApi) FindPetsByTagsGet(builder func(*PetApiFindPetsByTagsGetRequest)) *PetApiFindPetsByTagsGetRequest {
	result := &PetApiFindPetsByTagsGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Find pet by ID
// Returns a single pet
type PetApiGetPetByIdGetRequest struct {
	api *PetApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of pet to return
	PetId *int64
}

// Set PetId from non pointer
func (it *PetApiGetPetByIdGetRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	var model Pet
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *PetApiGetPetByIdGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet/{petId}")
	method := strings.ToUpper("Get")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
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
// petId: ID of pet to return
func (it *PetApi) GetPetByIdGet(builder func(*PetApiGetPetByIdGetRequest)) *PetApiGetPetByIdGetRequest {
	result := &PetApiGetPetByIdGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Update an existing pet
//
type PetApiUpdatePetPutRequest struct {
	api *PetApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// Pet object that needs to be added to the store
	Body *Pet
}

// Set Body from non pointer
func (it *PetApiUpdatePetPutRequest) SetBody(newBody Pet) {
	it.Body = &newBody
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
func (it *PetApiUpdatePetPutRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet")
	method := strings.ToUpper("Put")
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
// body: Pet object that needs to be added to the store
func (it *PetApi) UpdatePetPut(builder func(*PetApiUpdatePetPutRequest)) *PetApiUpdatePetPutRequest {
	result := &PetApiUpdatePetPutRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Updates a pet in the store with form data
//
type PetApiUpdatePetWithFormPostRequest struct {
	api *PetApi

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

// Set PetId from non pointer
func (it *PetApiUpdatePetWithFormPostRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Set Name from non pointer
func (it *PetApiUpdatePetWithFormPostRequest) SetName(newName string) {
	it.Name = &newName
}

// Set Status from non pointer
func (it *PetApiUpdatePetWithFormPostRequest) SetStatus(newStatus string) {
	it.Status = &newStatus
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
func (it *PetApiUpdatePetWithFormPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet/{petId}")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
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
// petId: ID of pet that needs to be updated
// name: Updated name of the pet
// status: Updated status of the pet
func (it *PetApi) UpdatePetWithFormPost(builder func(*PetApiUpdatePetWithFormPostRequest)) *PetApiUpdatePetWithFormPostRequest {
	result := &PetApiUpdatePetWithFormPostRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// uploads an image
//
type PetApiUploadFilePostRequest struct {
	api *PetApi

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

// Set PetId from non pointer
func (it *PetApiUploadFilePostRequest) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Set AdditionalMetadata from non pointer
func (it *PetApiUploadFilePostRequest) SetAdditionalMetadata(newAdditionalMetadata string) {
	it.AdditionalMetadata = &newAdditionalMetadata
}

// Set File from non pointer
func (it *PetApiUploadFilePostRequest) SetFile(newFile io.Reader) {
	it.File = &newFile
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
		return nil, nil, nil, xerrors.Errorf("http fetch failed, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, xerrors.Errorf("http read failed, %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, body, resp, xerrors.Errorf("http bad response, %w", err)
	}

	var model ApiResponse
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Fetch http request, returns raw response.
//
func (it *PetApiUploadFilePostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/pet/{petId}/uploadImage")
	method := strings.ToUpper("Post")
	var body io.Reader

	if it.PetId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"petId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.PetId)))
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
// petId: ID of pet to update
// additionalMetadata: Additional data to pass to server
// file: file to upload
func (it *PetApi) UploadFilePost(builder func(*PetApiUploadFilePostRequest)) *PetApiUploadFilePostRequest {
	result := &PetApiUploadFilePostRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

func (it *PetApi) this_is_call_dummy() {
	url.Parse("")
	xerrors.Errorf("")
	strings.ToUpper("")
	fmt.Sprintf("%v", "")
	io.ReadAtLeast(nil, nil, 0)
	json.Unmarshal(nil, nil)
}
