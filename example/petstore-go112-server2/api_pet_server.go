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

type PetApiController struct {
	router *mux.Router

	// Custom response write handler.
	Responder func(writer http.ResponseWriter, request *http.Request, handlerResponse SwaggerResponse, handlerError error)
}

// parse PetApiAddPetRequest object from http.Request
func ParsePetApiAddPetRequest(request *http.Request) (*PetApiAddPetPostRequest, error) {
	result := &PetApiAddPetPostRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return nil, xerrors.Errorf("Request body read failed, %w", err)
	} else {
		var model Pet
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
// Add a new pet to the store
//
//
func (it *PetApiController) HandleAddPetPost(handler func(httpRequest *http.Request, swaggerRequest *PetApiAddPetPostRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Post")).Path(path.Join("/v2", "/pet")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiAddPetRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse PetApiDeletePetRequest object from http.Request
func ParsePetApiDeletePetRequest(request *http.Request) (*PetApiDeletePetDeleteRequest, error) {
	result := &PetApiDeletePetDeleteRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "petId", &result.PetId)

	_ = parseValues(headerValues, "api_key", &result.ApiKey)

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
// Deletes a pet
//
//
func (it *PetApiController) HandleDeletePetDelete(handler func(httpRequest *http.Request, swaggerRequest *PetApiDeletePetDeleteRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Delete")).Path(path.Join("/v2", "/pet/{petId}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiDeletePetRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse PetApiFindPetsByStatusRequest object from http.Request
func ParsePetApiFindPetsByStatusRequest(request *http.Request) (*PetApiFindPetsByStatusGetRequest, error) {
	result := &PetApiFindPetsByStatusGetRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parseValues(queryValues, "status", &result.Status)

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
// Finds Pets by status
//
// Multiple status values can be provided with comma separated strings
func (it *PetApiController) HandleFindPetsByStatusGet(handler func(httpRequest *http.Request, swaggerRequest *PetApiFindPetsByStatusGetRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/pet/findByStatus")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiFindPetsByStatusRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse PetApiFindPetsByTagsRequest object from http.Request
func ParsePetApiFindPetsByTagsRequest(request *http.Request) (*PetApiFindPetsByTagsGetRequest, error) {
	result := &PetApiFindPetsByTagsGetRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parseValues(queryValues, "tags", &result.Tags)

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
// Finds Pets by tags
//
// Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
func (it *PetApiController) HandleFindPetsByTagsGet(handler func(httpRequest *http.Request, swaggerRequest *PetApiFindPetsByTagsGetRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/pet/findByTags")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiFindPetsByTagsRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse PetApiGetPetByIdRequest object from http.Request
func ParsePetApiGetPetByIdRequest(request *http.Request) (*PetApiGetPetByIdGetRequest, error) {
	result := &PetApiGetPetByIdGetRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "petId", &result.PetId)

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
// Find pet by ID
//
// Returns a single pet
func (it *PetApiController) HandleGetPetByIdGet(handler func(httpRequest *http.Request, swaggerRequest *PetApiGetPetByIdGetRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/pet/{petId}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiGetPetByIdRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse PetApiUpdatePetRequest object from http.Request
func ParsePetApiUpdatePetRequest(request *http.Request) (*PetApiUpdatePetPutRequest, error) {
	result := &PetApiUpdatePetPutRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return nil, xerrors.Errorf("Request body read failed, %w", err)
	} else {
		var model Pet
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
// Update an existing pet
//
//
func (it *PetApiController) HandleUpdatePetPut(handler func(httpRequest *http.Request, swaggerRequest *PetApiUpdatePetPutRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Put")).Path(path.Join("/v2", "/pet")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiUpdatePetRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse PetApiUpdatePetWithFormRequest object from http.Request
func ParsePetApiUpdatePetWithFormRequest(request *http.Request) (*PetApiUpdatePetWithFormPostRequest, error) {
	result := &PetApiUpdatePetWithFormPostRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "petId", &result.PetId)

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
// Updates a pet in the store with form data
//
//
func (it *PetApiController) HandleUpdatePetWithFormPost(handler func(httpRequest *http.Request, swaggerRequest *PetApiUpdatePetWithFormPostRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Post")).Path(path.Join("/v2", "/pet/{petId}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiUpdatePetWithFormRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse PetApiUploadFileRequest object from http.Request
func ParsePetApiUploadFileRequest(request *http.Request) (*PetApiUploadFilePostRequest, error) {
	result := &PetApiUploadFilePostRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "petId", &result.PetId)

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
// uploads an image
//
//
func (it *PetApiController) HandleUploadFilePost(handler func(httpRequest *http.Request, swaggerRequest *PetApiUploadFilePostRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Post")).Path(path.Join("/v2", "/pet/{petId}/uploadImage")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParsePetApiUploadFileRequest(request)
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
//  controller := petstore.NewPetApiController(router)
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
func NewPetApiController(router *mux.Router) *PetApiController {
	return &PetApiController{
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
func _PetApi_this_is_call_dummy() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}
