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

type AddPetPostParams struct {
	Request *http.Request
	// Pet object that needs to be added to the store
	Body *Pet
}

// Require value of Body
func (it *AddPetPostParams) RequireBody() Pet {
	if it.Body == nil {
		panic(xerrors.Errorf(".Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *AddPetPostParams) GetBody() Pet {
	if it.Body != nil {
		return *it.Body
	}
	result := new(Pet)
	return *result
}

func ParseAddPetPostParams(request *http.Request) (*AddPetPostParams, error) {
	result := &AddPetPostParams{
		Request: request,
	}

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

type DeletePetDeleteParams struct {
	Request *http.Request
	// Pet id to delete
	PetId *int64
	//
	ApiKey *string
}

// Require value of PetId
func (it *DeletePetDeleteParams) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf(".PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *DeletePetDeleteParams) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Require value of ApiKey
func (it *DeletePetDeleteParams) RequireApiKey() string {
	if it.ApiKey == nil {
		panic(xerrors.Errorf(".ApiKey is nil"))
	}
	return *it.ApiKey
}

// Get value of ApiKey / or default
func (it *DeletePetDeleteParams) GetApiKey() string {
	if it.ApiKey != nil {
		return *it.ApiKey
	}
	result := new(string)
	return *result
}

func ParseDeletePetDeleteParams(request *http.Request) (*DeletePetDeleteParams, error) {
	result := &DeletePetDeleteParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "petId", &result.PetId); err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	if err := parseValues(headerValues, "api_key", &result.ApiKey); err != nil {
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

type FindPetsByStatusGetParams struct {
	Request *http.Request
	// Status values that need to be considered for filter
	Status *[]string
}

// Require value of Status
func (it *FindPetsByStatusGetParams) RequireStatus() []string {
	if it.Status == nil {
		panic(xerrors.Errorf(".Status is nil"))
	}
	return *it.Status
}

// Get value of Status / or default
func (it *FindPetsByStatusGetParams) GetStatus() []string {
	if it.Status != nil {
		return *it.Status
	}
	result := new([]string)
	return *result
}

func ParseFindPetsByStatusGetParams(request *http.Request) (*FindPetsByStatusGetParams, error) {
	result := &FindPetsByStatusGetParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parseValues(queryValues, "status", &result.Status); err != nil {
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

type FindPetsByTagsGetParams struct {
	Request *http.Request
	// Tags to filter by
	Tags *[]string
}

// Require value of Tags
func (it *FindPetsByTagsGetParams) RequireTags() []string {
	if it.Tags == nil {
		panic(xerrors.Errorf(".Tags is nil"))
	}
	return *it.Tags
}

// Get value of Tags / or default
func (it *FindPetsByTagsGetParams) GetTags() []string {
	if it.Tags != nil {
		return *it.Tags
	}
	result := new([]string)
	return *result
}

func ParseFindPetsByTagsGetParams(request *http.Request) (*FindPetsByTagsGetParams, error) {
	result := &FindPetsByTagsGetParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parseValues(queryValues, "tags", &result.Tags); err != nil {
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

type GetPetByIdGetParams struct {
	Request *http.Request
	// ID of pet to return
	PetId *int64
}

// Require value of PetId
func (it *GetPetByIdGetParams) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf(".PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *GetPetByIdGetParams) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

func ParseGetPetByIdGetParams(request *http.Request) (*GetPetByIdGetParams, error) {
	result := &GetPetByIdGetParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "petId", &result.PetId); err != nil {
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

type UpdatePetPutParams struct {
	Request *http.Request
	// Pet object that needs to be added to the store
	Body *Pet
}

// Require value of Body
func (it *UpdatePetPutParams) RequireBody() Pet {
	if it.Body == nil {
		panic(xerrors.Errorf(".Body is nil"))
	}
	return *it.Body
}

// Get value of Body / or default
func (it *UpdatePetPutParams) GetBody() Pet {
	if it.Body != nil {
		return *it.Body
	}
	result := new(Pet)
	return *result
}

func ParseUpdatePetPutParams(request *http.Request) (*UpdatePetPutParams, error) {
	result := &UpdatePetPutParams{
		Request: request,
	}

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

type UpdatePetWithFormPostParams struct {
	Request *http.Request
	// ID of pet that needs to be updated
	PetId *int64
	// Updated name of the pet
	Name *string
	// Updated status of the pet
	Status *string
}

// Require value of PetId
func (it *UpdatePetWithFormPostParams) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf(".PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *UpdatePetWithFormPostParams) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Require value of Name
func (it *UpdatePetWithFormPostParams) RequireName() string {
	if it.Name == nil {
		panic(xerrors.Errorf(".Name is nil"))
	}
	return *it.Name
}

// Get value of Name / or default
func (it *UpdatePetWithFormPostParams) GetName() string {
	if it.Name != nil {
		return *it.Name
	}
	result := new(string)
	return *result
}

// Require value of Status
func (it *UpdatePetWithFormPostParams) RequireStatus() string {
	if it.Status == nil {
		panic(xerrors.Errorf(".Status is nil"))
	}
	return *it.Status
}

// Get value of Status / or default
func (it *UpdatePetWithFormPostParams) GetStatus() string {
	if it.Status != nil {
		return *it.Status
	}
	result := new(string)
	return *result
}

func ParseUpdatePetWithFormPostParams(request *http.Request) (*UpdatePetWithFormPostParams, error) {
	result := &UpdatePetWithFormPostParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "petId", &result.PetId); err != nil {
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

type UploadFilePostParams struct {
	Request *http.Request
	// ID of pet to update
	PetId *int64
	// Additional data to pass to server
	AdditionalMetadata *string
	// file to upload
	File *io.Reader
}

// Require value of PetId
func (it *UploadFilePostParams) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf(".PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *UploadFilePostParams) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Require value of AdditionalMetadata
func (it *UploadFilePostParams) RequireAdditionalMetadata() string {
	if it.AdditionalMetadata == nil {
		panic(xerrors.Errorf(".AdditionalMetadata is nil"))
	}
	return *it.AdditionalMetadata
}

// Get value of AdditionalMetadata / or default
func (it *UploadFilePostParams) GetAdditionalMetadata() string {
	if it.AdditionalMetadata != nil {
		return *it.AdditionalMetadata
	}
	result := new(string)
	return *result
}

// Require value of File
func (it *UploadFilePostParams) RequireFile() io.Reader {
	if it.File == nil {
		panic(xerrors.Errorf(".File is nil"))
	}
	return *it.File
}

// Get value of File / or default
func (it *UploadFilePostParams) GetFile() io.Reader {
	if it.File != nil {
		return *it.File
	}
	result := new(io.Reader)
	return *result
}

func ParseUploadFilePostParams(request *http.Request) (*UploadFilePostParams, error) {
	result := &UploadFilePostParams{
		Request: request,
	}

	pathValues := mux.Vars(request)
	queryValues := request.URL.Query()
	headerValues := request.Header

	if err := parsePath(pathValues, "petId", &result.PetId); err != nil {
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

type PetApiController struct {
	router *mux.Router

	// Custom response write handler.
	Responder func(writer http.ResponseWriter, request *http.Request, handlerResponse SwaggerResponse, handlerError error)
}

// Register handler to Router.
// Add a new pet to the store
//
//
func (it *PetApiController) HandleAddPetPost(handler func(request *http.Request, params *AddPetPostParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseAddPetPostParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Deletes a pet
//
//
func (it *PetApiController) HandleDeletePetDelete(handler func(request *http.Request, params *DeletePetDeleteParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseDeletePetDeleteParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Finds Pets by status
//
// Multiple status values can be provided with comma separated strings
func (it *PetApiController) HandleFindPetsByStatusGet(handler func(request *http.Request, params *FindPetsByStatusGetParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseFindPetsByStatusGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Finds Pets by tags
//
// Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
func (it *PetApiController) HandleFindPetsByTagsGet(handler func(request *http.Request, params *FindPetsByTagsGetParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseFindPetsByTagsGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Find pet by ID
//
// Returns a single pet
func (it *PetApiController) HandleGetPetByIdGet(handler func(request *http.Request, params *GetPetByIdGetParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseGetPetByIdGetParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Update an existing pet
//
//
func (it *PetApiController) HandleUpdatePetPut(handler func(request *http.Request, params *UpdatePetPutParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseUpdatePetPutParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// Updates a pet in the store with form data
//
//
func (it *PetApiController) HandleUpdatePetWithFormPost(handler func(request *http.Request, params *UpdatePetWithFormPostParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseUpdatePetWithFormPostParams(request)
		if err != nil {
			responseError = err
		} else {
			response, responseError = handler(request, params)
		}
	})
}

// Register handler to Router.
// uploads an image
//
//
func (it *PetApiController) HandleUploadFilePost(handler func(request *http.Request, params *UploadFilePostParams) (response SwaggerResponse, err error)) {
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

		params, err := ParseUploadFilePostParams(request)
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
