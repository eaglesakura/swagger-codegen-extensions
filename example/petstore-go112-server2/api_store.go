package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type DeleteOrderDeleteParams struct {
	Request *http.Request
	// ID of the order that needs to be deleted
	OrderId *int64
}

// Require value of OrderId
func (it *DeleteOrderDeleteParams) RequireOrderId() int64 {
	if it.OrderId == nil {
		panic(xerrors.Errorf(".OrderId is nil"))
	}
	return *it.OrderId
}

// Get value of OrderId / or default
func (it *DeleteOrderDeleteParams) GetOrderId() int64 {
	if it.OrderId != nil {
		return *it.OrderId
	}
	result := new(int64)
	return *result
}

func ParseDeleteOrderDeleteParams(request *http.Request) (*DeleteOrderDeleteParams, error) {
	result := &DeleteOrderDeleteParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "orderId", &result.OrderId); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	// dummy access for Compiler
	//noinspection GoSelfAssignment
	pathValues = pathValues
	//noinspection GoSelfAssignment
	queryValues = queryValues
	//noinspection GoSelfAssignment
	headerValues = headerValues

	return result, nil
}

type GetInventoryGetParams struct {
	Request *http.Request
}

func ParseGetInventoryGetParams(request *http.Request) (*GetInventoryGetParams, error) {
	result := &GetInventoryGetParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	// dummy access for Compiler
	//noinspection GoSelfAssignment
	pathValues = pathValues
	//noinspection GoSelfAssignment
	queryValues = queryValues
	//noinspection GoSelfAssignment
	headerValues = headerValues

	return result, nil
}

type GetOrderByIdGetParams struct {
	Request *http.Request
	// ID of pet that needs to be fetched
	OrderId *int64
}

// Require value of OrderId
func (it *GetOrderByIdGetParams) RequireOrderId() int64 {
	if it.OrderId == nil {
		panic(xerrors.Errorf(".OrderId is nil"))
	}
	return *it.OrderId
}

// Get value of OrderId / or default
func (it *GetOrderByIdGetParams) GetOrderId() int64 {
	if it.OrderId != nil {
		return *it.OrderId
	}
	result := new(int64)
	return *result
}

func ParseGetOrderByIdGetParams(request *http.Request) (*GetOrderByIdGetParams, error) {
	result := &GetOrderByIdGetParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "orderId", &result.OrderId); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	// dummy access for Compiler
	//noinspection GoSelfAssignment
	pathValues = pathValues
	//noinspection GoSelfAssignment
	queryValues = queryValues
	//noinspection GoSelfAssignment
	headerValues = headerValues

	return result, nil
}

type PlaceOrderPostParams struct {
	Request *http.Request
	// order placed for purchasing the pet
	Body *Order
}

// Require value of Body
func (it *PlaceOrderPostParams) RequireBody() Order {
	if it.Body == nil {
		panic(xerrors.Errorf(".Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *PlaceOrderPostParams) GetBody() Order {
	if it.Body != nil {
		return *it.Body
	}
	result := new(Order)
	return *result
}

func ParsePlaceOrderPostParams(request *http.Request) (*PlaceOrderPostParams, error) {
	result := &PlaceOrderPostParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return nil, xerrors.Errorf("Request body read failed, %w", err)
	} else {
		var model Order
		err = json.Unmarshal(body, &model)
		if err != nil {
			return nil, xerrors.Errorf("Request body parse failed, %w", err)
		}
		result.Body = &model
	}

	// dummy access for Compiler
	//noinspection GoSelfAssignment
	pathValues = pathValues
	//noinspection GoSelfAssignment
	queryValues = queryValues
	//noinspection GoSelfAssignment
	headerValues = headerValues

	return result, nil
}

type StoreApiController struct {
	router *mux.Router

	// Custom response write handler.
	Responder func(writer http.ResponseWriter, request *http.Request, handlerResponse SwaggerResponse, handlerError error)
}

// Register handler to Router.
// Delete purchase order by ID
//
// For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
func (it *StoreApiController) HandleDeleteOrderDelete(handler func(request *http.Request, params *DeleteOrderDeleteParams) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Delete")).Path(path.Join("/v2", "/store/order/{orderId}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseDeleteOrderDeleteParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Returns pet inventories by status
//
// Returns a map of status codes to quantities
func (it *StoreApiController) HandleGetInventoryGet(handler func(request *http.Request, params *GetInventoryGetParams) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/store/inventory")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseGetInventoryGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Find purchase order by ID
//
// For valid response try integer IDs with value >= 1 and <= 10.         Other values will generated exceptions
func (it *StoreApiController) HandleGetOrderByIdGet(handler func(request *http.Request, params *GetOrderByIdGetParams) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/store/order/{orderId}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseGetOrderByIdGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Place an order for a pet
//
//
func (it *StoreApiController) HandlePlaceOrderPost(handler func(request *http.Request, params *PlaceOrderPostParams) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Post")).Path(path.Join("/v2", "/store/order")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePlaceOrderPostParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// make controller.
//
//
// e.g.) How to use on HttpServer.
//  router := mux.NewRouter()
//  controller := petstore.NewStoreApiController(router)
//  controller.Handlers.HandlePathToFooApi(func(req, params) (SwaggerResponse, error) {
//      // Handle Foo API.
//  })
//  controller.Handlers.HandlePathToBarApi(func(req, params) (SwaggerResponse, error) {
//      // Handle Bar API.
//  })
//  http.Handle("/", router)
//
// e.g.) How to use on Google Cloud Functions
//  func FooCloudFunction(writer http.ResponseWriter, request *http.Request) {
//      router.ServeHTTP(writer, request)
//  }
func NewStoreApiController(router *mux.Router) *StoreApiController {
	return &StoreApiController{
		router: router,
		Responder: func(writer http.ResponseWriter, request *http.Request, handlerResponse SwaggerResponse, handlerError error) {
			if handlerError != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(handlerError.Error()))
				return
			}

			if handlerResponse != nil {
				handlerResponse.Write(writer, request)
				return
			}

			writer.WriteHeader(http.StatusNoContent)
		},
	}
}

//noinspection GoUnusedFunction,GoSnakeCaseUsage
func _StoreApi_this_is_call_dummy() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}
