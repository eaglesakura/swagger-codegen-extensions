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

type StoreApi struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

func NewStoreApi() *StoreApi {
	return &StoreApi{
		Endpoint: "",
	}
}

// Delete purchase order by ID
// For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
type StoreApiDeleteOrderDeleteRequest struct {
	api *StoreApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of the order that needs to be deleted
	OrderId *int64
}

// Set OrderId from non pointer
func (it *StoreApiDeleteOrderDeleteRequest) SetOrderId(newOrderId int64) {
	it.OrderId = &newOrderId
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

// Fetch http request, returns raw response.
//
func (it *StoreApiDeleteOrderDeleteRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/store/order/{orderId}")
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
// orderId: ID of the order that needs to be deleted
func (it *StoreApi) DeleteOrderDelete(builder func(*StoreApiDeleteOrderDeleteRequest)) *StoreApiDeleteOrderDeleteRequest {
	result := &StoreApiDeleteOrderDeleteRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Returns pet inventories by status
// Returns a map of status codes to quantities
type StoreApiGetInventoryGetRequest struct {
	api *StoreApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
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

// Fetch http request, returns raw response.
//
func (it *StoreApiGetInventoryGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/store/inventory")
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
func (it *StoreApi) GetInventoryGet(builder func(*StoreApiGetInventoryGetRequest)) *StoreApiGetInventoryGetRequest {
	result := &StoreApiGetInventoryGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Find purchase order by ID
// For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions
type StoreApiGetOrderByIdGetRequest struct {
	api *StoreApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// ID of pet that needs to be fetched
	OrderId *int64
}

// Set OrderId from non pointer
func (it *StoreApiGetOrderByIdGetRequest) SetOrderId(newOrderId int64) {
	it.OrderId = &newOrderId
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

// Fetch http request, returns raw response.
//
func (it *StoreApiGetOrderByIdGetRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/store/order/{orderId}")
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
// orderId: ID of pet that needs to be fetched
func (it *StoreApi) GetOrderByIdGet(builder func(*StoreApiGetOrderByIdGetRequest)) *StoreApiGetOrderByIdGetRequest {
	result := &StoreApiGetOrderByIdGetRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

// Place an order for a pet
//
type StoreApiPlaceOrderPostRequest struct {
	api *StoreApi

	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)

	// order placed for purchasing the pet
	Body *Order
}

// Set Body from non pointer
func (it *StoreApiPlaceOrderPostRequest) SetBody(newBody Order) {
	it.Body = &newBody
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

// Fetch http request, returns raw response.
//
func (it *StoreApiPlaceOrderPostRequest) Fetch(ctx context.Context) (*http.Response, error) {
	client := http.DefaultClient
	apiUrl := path.Join(it.Endpoint, "/v2", "/store/order")
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
// body: order placed for purchasing the pet
func (it *StoreApi) PlaceOrderPost(builder func(*StoreApiPlaceOrderPostRequest)) *StoreApiPlaceOrderPostRequest {
	result := &StoreApiPlaceOrderPostRequest{
		api:       it,
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}

func (it *StoreApi) this_is_call_dummy() {
	url.Parse("")
	xerrors.Errorf("")
	strings.ToUpper("")
	fmt.Sprintf("%v", "")
	io.ReadAtLeast(nil, nil, 0)
	json.Unmarshal(nil, nil)
}
