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

type CreateUserPostParams struct {
	Request *http.Request
	// Created user object
	Body *User
}

// Require value of Body
func (it *CreateUserPostParams) RequireBody() User {
	if it.Body == nil {
		panic(xerrors.Errorf(".Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *CreateUserPostParams) GetBody() User {
	if it.Body != nil {
		return *it.Body
	}
	result := new(User)
	return *result
}

func ParseCreateUserPostParams(request *http.Request) (*CreateUserPostParams, error) {
	result := &CreateUserPostParams{
		Request: request,
	}

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

type CreateUsersWithArrayInputPostParams struct {
	Request *http.Request
	// List of user object
	Body *[]User
}

// Require value of Body
func (it *CreateUsersWithArrayInputPostParams) RequireBody() []User {
	if it.Body == nil {
		panic(xerrors.Errorf(".Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *CreateUsersWithArrayInputPostParams) GetBody() []User {
	if it.Body != nil {
		return *it.Body
	}
	result := new([]User)
	return *result
}

func ParseCreateUsersWithArrayInputPostParams(request *http.Request) (*CreateUsersWithArrayInputPostParams, error) {
	result := &CreateUsersWithArrayInputPostParams{
		Request: request,
	}

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

type CreateUsersWithListInputPostParams struct {
	Request *http.Request
	// List of user object
	Body *[]User
}

// Require value of Body
func (it *CreateUsersWithListInputPostParams) RequireBody() []User {
	if it.Body == nil {
		panic(xerrors.Errorf(".Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *CreateUsersWithListInputPostParams) GetBody() []User {
	if it.Body != nil {
		return *it.Body
	}
	result := new([]User)
	return *result
}

func ParseCreateUsersWithListInputPostParams(request *http.Request) (*CreateUsersWithListInputPostParams, error) {
	result := &CreateUsersWithListInputPostParams{
		Request: request,
	}

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

type DeleteUserDeleteParams struct {
	Request *http.Request
	// The name that needs to be deleted
	Username *string
}

// Require value of Username
func (it *DeleteUserDeleteParams) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf(".Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *DeleteUserDeleteParams) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

func ParseDeleteUserDeleteParams(request *http.Request) (*DeleteUserDeleteParams, error) {
	result := &DeleteUserDeleteParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "username", &result.Username); err != nil {
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

type GetUserByNameGetParams struct {
	Request *http.Request
	// The name that needs to be fetched. Use user1 for testing.
	Username *string
}

// Require value of Username
func (it *GetUserByNameGetParams) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf(".Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *GetUserByNameGetParams) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

func ParseGetUserByNameGetParams(request *http.Request) (*GetUserByNameGetParams, error) {
	result := &GetUserByNameGetParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "username", &result.Username); err != nil {
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

type LoginUserGetParams struct {
	Request *http.Request
	// The user name for login
	Username *string
	// The password for login in clear text
	Password *string
}

// Require value of Username
func (it *LoginUserGetParams) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf(".Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *LoginUserGetParams) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

// Require value of Password
func (it *LoginUserGetParams) RequirePassword() string {
	if it.Password == nil {
		panic(xerrors.Errorf(".Password is nil"))
	}
	return *it.Password
}

// Get value of Password / or default
func (it *LoginUserGetParams) GetPassword() string {
	if it.Password != nil {
		return *it.Password
	}
	result := new(string)
	return *result
}

func ParseLoginUserGetParams(request *http.Request) (*LoginUserGetParams, error) {
	result := &LoginUserGetParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parseValues(queryValues, "username", &result.Username); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	if err := parseValues(queryValues, "password", &result.Password); err != nil {
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

type LogoutUserGetParams struct {
	Request *http.Request
}

func ParseLogoutUserGetParams(request *http.Request) (*LogoutUserGetParams, error) {
	result := &LogoutUserGetParams{
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

type UpdateUserPutParams struct {
	Request *http.Request
	// name that need to be updated
	Username *string
	// Updated user object
	Body *User
}

// Require value of Username
func (it *UpdateUserPutParams) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf(".Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *UpdateUserPutParams) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

// Require value of Body
func (it *UpdateUserPutParams) RequireBody() User {
	if it.Body == nil {
		panic(xerrors.Errorf(".Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *UpdateUserPutParams) GetBody() User {
	if it.Body != nil {
		return *it.Body
	}
	result := new(User)
	return *result
}

func ParseUpdateUserPutParams(request *http.Request) (*UpdateUserPutParams, error) {
	result := &UpdateUserPutParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "username", &result.Username); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

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

type UserApiController struct {
	router *mux.Router

	// Custom response write handler.
	Responder func(writer http.ResponseWriter, request *http.Request, handlerResponse SwaggerResponse, handlerError error)
}

// Register handler to Router.
// Create user
//
// This can only be done by the logged in user.
func (it *UserApiController) HandleCreateUserPost(handler func(request *http.Request, params *CreateUserPostParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseCreateUserPostParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Creates list of users with given input array
//
//
func (it *UserApiController) HandleCreateUsersWithArrayInputPost(handler func(request *http.Request, params *CreateUsersWithArrayInputPostParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseCreateUsersWithArrayInputPostParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Creates list of users with given input array
//
//
func (it *UserApiController) HandleCreateUsersWithListInputPost(handler func(request *http.Request, params *CreateUsersWithListInputPostParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseCreateUsersWithListInputPostParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Delete user
//
// This can only be done by the logged in user.
func (it *UserApiController) HandleDeleteUserDelete(handler func(request *http.Request, params *DeleteUserDeleteParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseDeleteUserDeleteParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Get user by user name
//
//
func (it *UserApiController) HandleGetUserByNameGet(handler func(request *http.Request, params *GetUserByNameGetParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseGetUserByNameGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Logs user into the system
//
//
func (it *UserApiController) HandleLoginUserGet(handler func(request *http.Request, params *LoginUserGetParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseLoginUserGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Logs out current logged in user session
//
//
func (it *UserApiController) HandleLogoutUserGet(handler func(request *http.Request, params *LogoutUserGetParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseLogoutUserGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Updated user
//
// This can only be done by the logged in user.
func (it *UserApiController) HandleUpdateUserPut(handler func(request *http.Request, params *UpdateUserPutParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseUpdateUserPutParams(request)
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
