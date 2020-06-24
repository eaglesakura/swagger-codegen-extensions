package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"net/http"
)

// petstore
type StoreApi struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

// Update self from builder callback.
func (it *StoreApi) Update(builder func(ref *StoreApi)) *StoreApi {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *StoreApi) Copy(builder func(copied *StoreApi)) *StoreApi {
	result := *it
	builder(&result)
	return &result
}

func NewStoreApi() *StoreApi {
	return &StoreApi{
		Endpoint: "",
	}
}

func (it *StoreApi) DeleteOrderDelete(builder func(*StoreApiDeleteOrderDeleteRequest)) *StoreApiDeleteOrderDeleteRequest {
	result := &StoreApiDeleteOrderDeleteRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *StoreApi) GetInventoryGet(builder func(*StoreApiGetInventoryGetRequest)) *StoreApiGetInventoryGetRequest {
	result := &StoreApiGetInventoryGetRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *StoreApi) GetOrderByIdGet(builder func(*StoreApiGetOrderByIdGetRequest)) *StoreApiGetOrderByIdGetRequest {
	result := &StoreApiGetOrderByIdGetRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *StoreApi) PlaceOrderPost(builder func(*StoreApiPlaceOrderPostRequest)) *StoreApiPlaceOrderPostRequest {
	result := &StoreApiPlaceOrderPostRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
