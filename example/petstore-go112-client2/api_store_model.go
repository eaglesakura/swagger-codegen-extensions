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

// Delete purchase order by ID
// For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
type StoreApiDeleteOrderDeleteRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of the order that needs to be deleted
	OrderId *int64
}

// Validation parameter
func (it *StoreApiDeleteOrderDeleteRequest) Valid() error {

	if it.OrderId == nil {
		return xerrors.Errorf("required parameter(OrderId) read failed")
	}

	return nil
}

// Set OrderId from non pointer
func (it *StoreApiDeleteOrderDeleteRequest) SetOrderId(newOrderId int64) {
	it.OrderId = &newOrderId
}

// Remove OrderId property.
func (it *StoreApiDeleteOrderDeleteRequest) RemoveOrderId() {
	it.OrderId = nil
}

// Require value of OrderId
func (it *StoreApiDeleteOrderDeleteRequest) RequireOrderId() int64 {
	if it.OrderId == nil {
		panic(xerrors.Errorf("StoreApi.OrderId is nil"))
	}
	return *it.OrderId
}

// Get value of OrderId / or default
func (it *StoreApiDeleteOrderDeleteRequest) GetOrderId() int64 {
	if it.OrderId != nil {
		return *it.OrderId
	}
	result := new(int64)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *StoreApiDeleteOrderDeleteRequest) Execute(ctx context.Context) (success []byte, responseBody []byte, resp *http.Response, err error) {
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

// Build http request
func (it *StoreApiDeleteOrderDeleteRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/store/order/{orderId}")
	method := strings.ToUpper("Delete")
	var body io.Reader

	if it.OrderId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"orderId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.OrderId)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *StoreApiDeleteOrderDeleteRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed, %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *StoreApiDeleteOrderDeleteRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Returns pet inventories by status
// Returns a map of status codes to quantities
type StoreApiGetInventoryGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

// Validation parameter
func (it *StoreApiGetInventoryGetRequest) Valid() error {

	return nil
}

// Fetch http request, returns raw response.
//
func (it *StoreApiGetInventoryGetRequest) Execute(ctx context.Context) (success *map[string]int32, responseBody []byte, resp *http.Response, err error) {
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

	var model map[string]int32
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *StoreApiGetInventoryGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/store/inventory")
	method := strings.ToUpper("Get")
	var body io.Reader

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *StoreApiGetInventoryGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed, %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *StoreApiGetInventoryGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Find purchase order by ID
// For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions
type StoreApiGetOrderByIdGetRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of pet that needs to be fetched
	OrderId *int64
}

// Validation parameter
func (it *StoreApiGetOrderByIdGetRequest) Valid() error {

	if it.OrderId == nil {
		return xerrors.Errorf("required parameter(OrderId) read failed")
	}

	return nil
}

// Set OrderId from non pointer
func (it *StoreApiGetOrderByIdGetRequest) SetOrderId(newOrderId int64) {
	it.OrderId = &newOrderId
}

// Remove OrderId property.
func (it *StoreApiGetOrderByIdGetRequest) RemoveOrderId() {
	it.OrderId = nil
}

// Require value of OrderId
func (it *StoreApiGetOrderByIdGetRequest) RequireOrderId() int64 {
	if it.OrderId == nil {
		panic(xerrors.Errorf("StoreApi.OrderId is nil"))
	}
	return *it.OrderId
}

// Get value of OrderId / or default
func (it *StoreApiGetOrderByIdGetRequest) GetOrderId() int64 {
	if it.OrderId != nil {
		return *it.OrderId
	}
	result := new(int64)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *StoreApiGetOrderByIdGetRequest) Execute(ctx context.Context) (success *Order, responseBody []byte, resp *http.Response, err error) {
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

	var model Order
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *StoreApiGetOrderByIdGetRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/store/order/{orderId}")
	method := strings.ToUpper("Get")
	var body io.Reader

	if it.OrderId != nil {
		apiUrl = strings.ReplaceAll(apiUrl, "{"+"orderId"+"}", url.PathEscape(fmt.Sprintf("%v", *it.OrderId)))
	}

	request, err := http.NewRequest(method, apiUrl, body)
	if err != nil {
		return nil, xerrors.Errorf("create request failed, %w", err)
	}

	query := request.URL.Query()

	request.URL.RawQuery = query.Encode()
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *StoreApiGetOrderByIdGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed, %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *StoreApiGetOrderByIdGetRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}

// Place an order for a pet
//
type StoreApiPlaceOrderPostRequest struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// order placed for purchasing the pet
	Body *Order
}

// Validation parameter
func (it *StoreApiPlaceOrderPostRequest) Valid() error {

	if it.Body == nil {
		return xerrors.Errorf("required parameter(Body) read failed")
	}

	return nil
}

// Set Body from non pointer
func (it *StoreApiPlaceOrderPostRequest) SetBody(newBody Order) {
	it.Body = &newBody
}

// Remove Body property.
func (it *StoreApiPlaceOrderPostRequest) RemoveBody() {
	it.Body = nil
}

// Require value of Body
func (it *StoreApiPlaceOrderPostRequest) RequireBody() Order {
	if it.Body == nil {
		panic(xerrors.Errorf("StoreApi.Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *StoreApiPlaceOrderPostRequest) GetBody() Order {
	if it.Body != nil {
		return *it.Body
	}
	result := new(Order)
	return *result
}

// Fetch http request, returns raw response.
//
func (it *StoreApiPlaceOrderPostRequest) Execute(ctx context.Context) (success *Order, responseBody []byte, resp *http.Response, err error) {
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

	var model Order
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, body, resp, xerrors.Errorf("json parse failed, %w", err)
	}
	return &model, body, resp, nil
}

// Build http request
func (it *StoreApiPlaceOrderPostRequest) BuildHttpRequest() (*http.Request, error) {
	endpoint := it.Endpoint
	if len(endpoint) == 0 {
		endpoint = "http://127.0.0.1/"
	}
	apiUrl := path.Join(endpoint, "/v2", "/store/order")
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
	return request, nil
}

// Fetch http request, returns raw response.
//
func (it *StoreApiPlaceOrderPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	request, err := it.BuildHttpRequest()
	if err != nil {
		return nil, xerrors.Errorf("http request builed failed, %w", err)
	}
	if it.Intercept != nil {
		client, request = it.Intercept(client, request)
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, xerrors.Errorf("http fetch failed, %w", err)
	}
	return response, nil
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func (it *StoreApiPlaceOrderPostRequest) dummyForCompiler() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}
