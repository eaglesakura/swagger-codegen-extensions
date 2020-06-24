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

type UserApiController struct {
	router *mux.Router

	// Custom response write handler.
	Responder func(writer http.ResponseWriter, request *http.Request, handlerResponse SwaggerResponse, handlerError error)
}

// parse UserApiCreateUserRequest object from http.Request
func ParseUserApiCreateUserRequest(request *http.Request) (*UserApiCreateUserPostRequest, error) {
	result := &UserApiCreateUserPostRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return nil, xerrors.Errorf("Request body read failed, %w", err)
	} else {
		var model User
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
// Create user
//
// This can only be done by the logged in user.
func (it *UserApiController) HandleCreateUserPost(handler func(httpRequest *http.Request, swaggerRequest *UserApiCreateUserPostRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Post")).Path(path.Join("/v2", "/user")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiCreateUserRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse UserApiCreateUsersWithArrayInputRequest object from http.Request
func ParseUserApiCreateUsersWithArrayInputRequest(request *http.Request) (*UserApiCreateUsersWithArrayInputPostRequest, error) {
	result := &UserApiCreateUsersWithArrayInputPostRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return nil, xerrors.Errorf("Request body read failed, %w", err)
	} else {
		var model []User
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
// Creates list of users with given input array
//
//
func (it *UserApiController) HandleCreateUsersWithArrayInputPost(handler func(httpRequest *http.Request, swaggerRequest *UserApiCreateUsersWithArrayInputPostRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Post")).Path(path.Join("/v2", "/user/createWithArray")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiCreateUsersWithArrayInputRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse UserApiCreateUsersWithListInputRequest object from http.Request
func ParseUserApiCreateUsersWithListInputRequest(request *http.Request) (*UserApiCreateUsersWithListInputPostRequest, error) {
	result := &UserApiCreateUsersWithListInputPostRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return nil, xerrors.Errorf("Request body read failed, %w", err)
	} else {
		var model []User
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
// Creates list of users with given input array
//
//
func (it *UserApiController) HandleCreateUsersWithListInputPost(handler func(httpRequest *http.Request, swaggerRequest *UserApiCreateUsersWithListInputPostRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Post")).Path(path.Join("/v2", "/user/createWithList")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiCreateUsersWithListInputRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse UserApiDeleteUserRequest object from http.Request
func ParseUserApiDeleteUserRequest(request *http.Request) (*UserApiDeleteUserDeleteRequest, error) {
	result := &UserApiDeleteUserDeleteRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "username", &result.Username)

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
// Delete user
//
// This can only be done by the logged in user.
func (it *UserApiController) HandleDeleteUserDelete(handler func(httpRequest *http.Request, swaggerRequest *UserApiDeleteUserDeleteRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Delete")).Path(path.Join("/v2", "/user/{username}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiDeleteUserRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse UserApiGetUserByNameRequest object from http.Request
func ParseUserApiGetUserByNameRequest(request *http.Request) (*UserApiGetUserByNameGetRequest, error) {
	result := &UserApiGetUserByNameGetRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "username", &result.Username)

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
// Get user by user name
//
//
func (it *UserApiController) HandleGetUserByNameGet(handler func(httpRequest *http.Request, swaggerRequest *UserApiGetUserByNameGetRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/user/{username}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiGetUserByNameRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse UserApiLoginUserRequest object from http.Request
func ParseUserApiLoginUserRequest(request *http.Request) (*UserApiLoginUserGetRequest, error) {
	result := &UserApiLoginUserGetRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parseValues(queryValues, "username", &result.Username)
	_ = parseValues(queryValues, "password", &result.Password)

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
// Logs user into the system
//
//
func (it *UserApiController) HandleLoginUserGet(handler func(httpRequest *http.Request, swaggerRequest *UserApiLoginUserGetRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/user/login")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiLoginUserRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse UserApiLogoutUserRequest object from http.Request
func ParseUserApiLogoutUserRequest(request *http.Request) (*UserApiLogoutUserGetRequest, error) {
	result := &UserApiLogoutUserGetRequest{}

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
// Logs out current logged in user session
//
//
func (it *UserApiController) HandleLogoutUserGet(handler func(httpRequest *http.Request, swaggerRequest *UserApiLogoutUserGetRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Get")).Path(path.Join("/v2", "/user/logout")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiLogoutUserRequest(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// parse UserApiUpdateUserRequest object from http.Request
func ParseUserApiUpdateUserRequest(request *http.Request) (*UserApiUpdateUserPutRequest, error) {
	result := &UserApiUpdateUserPutRequest{}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	_ = parsePath(pathValues, "username", &result.Username)

	if body, err := ioutil.ReadAll(request.Body); err != nil {
		return nil, xerrors.Errorf("Request body read failed, %w", err)
	} else {
		var model User
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
// Updated user
//
// This can only be done by the logged in user.
func (it *UserApiController) HandleUpdateUserPut(handler func(httpRequest *http.Request, swaggerRequest *UserApiUpdateUserPutRequest) (response SwaggerResponse, err error)) {
	r := it.router
	r.Methods(strings.ToUpper("Put")).Path(path.Join("/v2", "/user/{username}")).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var response SwaggerResponse
		var responseError error
		defer func() {
			recoverError := recover()
			if err, ok := recoverError.(error); ok {
				responseError = err
			}

			it.Responder(writer, request, response, responseError)
		}()

		params, err := ParseUserApiUpdateUserRequest(request)
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
//  controller := petstore.NewUserApiController(router)
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
func NewUserApiController(router *mux.Router) *UserApiController {
	return &UserApiController{
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
func _UserApi_this_is_call_dummy() {
	_, _ = ioutil.ReadAll(nil)
	_, _ = url.Parse("")
	_ = xerrors.Errorf("")
	strings.ToUpper("")
	_ = fmt.Sprintf("%v", "")
	_, _ = io.ReadAtLeast(nil, nil, 0)
	_ = json.Unmarshal(nil, nil)
}
