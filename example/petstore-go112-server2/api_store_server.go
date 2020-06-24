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

type StoreApiController struct {
	router *mux.Router

	// Custom response write handler.
	Responder func(writer http.ResponseWriter, request *http.Request, handlerResponse SwaggerResponse, handlerError error)
}

// parse StoreApiDeleteOrderRequest object from http.Request
func ParseStoreApiDeleteOrderRequest(request *http.Request) (*StoreApiDeleteOrderDeleteRequest, error) {
	result := &StoreApiDeleteOrderDeleteRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "orderId", &result.OrderId)

	// dummy access for Compiler
	//noinspection GoSelfAssignment
	pathValues = pathValues
	//noinspection GoSelfAssignment
	queryValues = queryValues
	//noinspection GoSelfAssignment
	headerValues = headerValues

	return result, nil
}

// Register handler to Router.
// Delete purchase order by ID
//
// For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
func (it *StoreApiController) HandleDeleteOrderDelete(handler func(httpRequest *http.Request, swaggerRequest *StoreApiDeleteOrderDeleteRequest) (response SwaggerResponse, err error)) {
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

		params, err := ParseStoreApiDeleteOrderRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse StoreApiGetInventoryRequest object from http.Request
func ParseStoreApiGetInventoryRequest(request *http.Request) (*StoreApiGetInventoryGetRequest, error) {
	result := &StoreApiGetInventoryGetRequest{}

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

// Register handler to Router.
// Returns pet inventories by status
//
// Returns a map of status codes to quantities
func (it *StoreApiController) HandleGetInventoryGet(handler func(httpRequest *http.Request, swaggerRequest *StoreApiGetInventoryGetRequest) (response SwaggerResponse, err error)) {
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

		params, err := ParseStoreApiGetInventoryRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse StoreApiGetOrderByIdRequest object from http.Request
func ParseStoreApiGetOrderByIdRequest(request *http.Request) (*StoreApiGetOrderByIdGetRequest, error) {
	result := &StoreApiGetOrderByIdGetRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "orderId", &result.OrderId)

	// dummy access for Compiler
	//noinspection GoSelfAssignment
	pathValues = pathValues
	//noinspection GoSelfAssignment
	queryValues = queryValues
	//noinspection GoSelfAssignment
	headerValues = headerValues

	return result, nil
}

// Register handler to Router.
// Find purchase order by ID
//
// For valid response try integer IDs with value >= 1 and <= 10.         Other values will generated exceptions
func (it *StoreApiController) HandleGetOrderByIdGet(handler func(httpRequest *http.Request, swaggerRequest *StoreApiGetOrderByIdGetRequest) (response SwaggerResponse, err error)) {
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

		params, err := ParseStoreApiGetOrderByIdRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse StoreApiPlaceOrderRequest object from http.Request
func ParseStoreApiPlaceOrderRequest(request *http.Request) (*StoreApiPlaceOrderPostRequest, error) {
	result := &StoreApiPlaceOrderPostRequest{}

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

// Register handler to Router.
// Place an order for a pet
//
//
func (it *StoreApiController) HandlePlaceOrderPost(handler func(httpRequest *http.Request, swaggerRequest *StoreApiPlaceOrderPostRequest) (response SwaggerResponse, err error)) {
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

		params, err := ParseStoreApiPlaceOrderRequest(request)
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
