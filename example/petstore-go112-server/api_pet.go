package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"net/http"
	"strings"
)

type AddPetParams struct {
	// Pet object that needs to be added to the store
	Body *Pet
}

/*
Add a new pet to the store


 param: Body Pet object that needs to be added to the store
 return: void
*/
type AddPetHandler func(context swagger.RequestContext, params *AddPetParams) swagger.Responder

func (it *AddPetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Body, it.Body == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewAddPetParams(binder swagger.RequestBinder) (*AddPetParams, error) {
	result := &AddPetParams{}

	if err := binder.BindBody("Pet", &result.Body); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type DeletePetParams struct {
	// Pet id to delete
	PetId *int64
	//
	ApiKey *string
}

/*
Deletes a pet


 param: PetId Pet id to delete
 param: ApiKey
 return: void
*/
type DeletePetHandler func(context swagger.RequestContext, params *DeletePetParams) swagger.Responder

func (it *DeletePetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.PetId, it.PetId == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.ApiKey, it.ApiKey == nil).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewDeletePetParams(binder swagger.RequestBinder) (*DeletePetParams, error) {
	result := &DeletePetParams{}

	if err := binder.BindPath("petId", "int64", &result.PetId); err != nil {
		return nil, err
	}

	if err := binder.BindHeader("api_key", "string", &result.ApiKey); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type FindPetsByStatusParams struct {
	// Status values that need to be considered for filter
	Status *[]string
}

/*
Finds Pets by status

Multiple status values can be provided with comma separated strings
 param: Status Status values that need to be considered for filter
 return: []Pet
*/
type FindPetsByStatusHandler func(context swagger.RequestContext, params *FindPetsByStatusParams) swagger.Responder

func (it *FindPetsByStatusParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Status, it.Status == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewFindPetsByStatusParams(binder swagger.RequestBinder) (*FindPetsByStatusParams, error) {
	result := &FindPetsByStatusParams{}

	if err := binder.BindQuery("status", "[]string", &result.Status); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type FindPetsByTagsParams struct {
	// Tags to filter by
	Tags *[]string
}

/*
Finds Pets by tags

Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
 param: Tags Tags to filter by
 return: []Pet
*/
type FindPetsByTagsHandler func(context swagger.RequestContext, params *FindPetsByTagsParams) swagger.Responder

func (it *FindPetsByTagsParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Tags, it.Tags == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewFindPetsByTagsParams(binder swagger.RequestBinder) (*FindPetsByTagsParams, error) {
	result := &FindPetsByTagsParams{}

	if err := binder.BindQuery("tags", "[]string", &result.Tags); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type GetPetByIdParams struct {
	// ID of pet to return
	PetId *int64
}

/*
Find pet by ID

Returns a single pet
 param: PetId ID of pet to return
 return: Pet
*/
type GetPetByIdHandler func(context swagger.RequestContext, params *GetPetByIdParams) swagger.Responder

func (it *GetPetByIdParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.PetId, it.PetId == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewGetPetByIdParams(binder swagger.RequestBinder) (*GetPetByIdParams, error) {
	result := &GetPetByIdParams{}

	if err := binder.BindPath("petId", "int64", &result.PetId); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type UpdatePetParams struct {
	// Pet object that needs to be added to the store
	Body *Pet
}

/*
Update an existing pet


 param: Body Pet object that needs to be added to the store
 return: void
*/
type UpdatePetHandler func(context swagger.RequestContext, params *UpdatePetParams) swagger.Responder

func (it *UpdatePetParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Body, it.Body == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewUpdatePetParams(binder swagger.RequestBinder) (*UpdatePetParams, error) {
	result := &UpdatePetParams{}

	if err := binder.BindBody("Pet", &result.Body); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type UpdatePetWithFormParams struct {
	// ID of pet that needs to be updated
	PetId *int64
	// Updated name of the pet
	Name *string
	// Updated status of the pet
	Status *string
}

/*
Updates a pet in the store with form data


 param: PetId ID of pet that needs to be updated
 param: Name Updated name of the pet
 param: Status Updated status of the pet
 return: void
*/
type UpdatePetWithFormHandler func(context swagger.RequestContext, params *UpdatePetWithFormParams) swagger.Responder

func (it *UpdatePetWithFormParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.PetId, it.PetId == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Name, it.Name == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Status, it.Status == nil).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewUpdatePetWithFormParams(binder swagger.RequestBinder) (*UpdatePetWithFormParams, error) {
	result := &UpdatePetWithFormParams{}

	if err := binder.BindPath("petId", "int64", &result.PetId); err != nil {
		return nil, err
	}

	if err := binder.BindForm("name", "string", &result.Name); err != nil {
		return nil, err
	}
	if err := binder.BindForm("status", "string", &result.Status); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type UploadFileParams struct {
	// ID of pet to update
	PetId *int64
	// Additional data to pass to server
	AdditionalMetadata *string
	// file to upload
	File *io.Reader
}

/*
uploads an image


 param: PetId ID of pet to update
 param: AdditionalMetadata Additional data to pass to server
 param: File file to upload
 return: ApiResponse
*/
type UploadFileHandler func(context swagger.RequestContext, params *UploadFileParams) swagger.Responder

func (it *UploadFileParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.PetId, it.PetId == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.AdditionalMetadata, it.AdditionalMetadata == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.File, it.File == nil).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewUploadFileParams(binder swagger.RequestBinder) (*UploadFileParams, error) {
	result := &UploadFileParams{}

	if err := binder.BindPath("petId", "int64", &result.PetId); err != nil {
		return nil, err
	}

	if err := binder.BindForm("additionalMetadata", "string", &result.AdditionalMetadata); err != nil {
		return nil, err
	}
	if err := binder.BindForm("file", "io.Reader", &result.File); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type PetApiController struct {
	AddPet swagger.HandleRequest

	DeletePet swagger.HandleRequest

	FindPetsByStatus swagger.HandleRequest

	FindPetsByTags swagger.HandleRequest

	GetPetById swagger.HandleRequest

	UpdatePet swagger.HandleRequest

	UpdatePetWithForm swagger.HandleRequest

	UploadFile swagger.HandleRequest
}

func NewPetApiController() *PetApiController {
	result := &PetApiController{}

	result.AddPet.Path = "/v2/pet"
	result.AddPet.Method = strings.ToUpper("Post")
	result.HandleAddPet(func(context swagger.RequestContext, params *AddPetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl AddPet"))
	})

	result.DeletePet.Path = "/v2/pet/{petId}"
	result.DeletePet.Method = strings.ToUpper("Delete")
	result.HandleDeletePet(func(context swagger.RequestContext, params *DeletePetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl DeletePet"))
	})

	result.FindPetsByStatus.Path = "/v2/pet/findByStatus"
	result.FindPetsByStatus.Method = strings.ToUpper("Get")
	result.HandleFindPetsByStatus(func(context swagger.RequestContext, params *FindPetsByStatusParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl FindPetsByStatus"))
	})

	result.FindPetsByTags.Path = "/v2/pet/findByTags"
	result.FindPetsByTags.Method = strings.ToUpper("Get")
	result.HandleFindPetsByTags(func(context swagger.RequestContext, params *FindPetsByTagsParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl FindPetsByTags"))
	})

	result.GetPetById.Path = "/v2/pet/{petId}"
	result.GetPetById.Method = strings.ToUpper("Get")
	result.HandleGetPetById(func(context swagger.RequestContext, params *GetPetByIdParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl GetPetById"))
	})

	result.UpdatePet.Path = "/v2/pet"
	result.UpdatePet.Method = strings.ToUpper("Put")
	result.HandleUpdatePet(func(context swagger.RequestContext, params *UpdatePetParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl UpdatePet"))
	})

	result.UpdatePetWithForm.Path = "/v2/pet/{petId}"
	result.UpdatePetWithForm.Method = strings.ToUpper("Post")
	result.HandleUpdatePetWithForm(func(context swagger.RequestContext, params *UpdatePetWithFormParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl UpdatePetWithForm"))
	})

	result.UploadFile.Path = "/v2/pet/{petId}/uploadImage"
	result.UploadFile.Method = strings.ToUpper("Post")
	result.HandleUploadFile(func(context swagger.RequestContext, params *UploadFileParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl UploadFile"))
	})

	return result
}

func (it *PetApiController) HandleAddPet(handler AddPetHandler) {
	it.AddPet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewAddPetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) HandleDeletePet(handler DeletePetHandler) {
	it.DeletePet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewDeletePetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) HandleFindPetsByStatus(handler FindPetsByStatusHandler) {
	it.FindPetsByStatus.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewFindPetsByStatusParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) HandleFindPetsByTags(handler FindPetsByTagsHandler) {
	it.FindPetsByTags.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewFindPetsByTagsParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) HandleGetPetById(handler GetPetByIdHandler) {
	it.GetPetById.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewGetPetByIdParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) HandleUpdatePet(handler UpdatePetHandler) {
	it.UpdatePet.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewUpdatePetParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) HandleUpdatePetWithForm(handler UpdatePetWithFormHandler) {
	it.UpdatePetWithForm.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewUpdatePetWithFormParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) HandleUploadFile(handler UploadFileHandler) {
	it.UploadFile.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewUploadFileParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *PetApiController) MapHandlers(mapper swagger.HandleMapper) {

	mapper.PutHandler(it.AddPet)

	mapper.PutHandler(it.DeletePet)

	mapper.PutHandler(it.FindPetsByStatus)

	mapper.PutHandler(it.FindPetsByTags)

	mapper.PutHandler(it.GetPetById)

	mapper.PutHandler(it.UpdatePet)

	mapper.PutHandler(it.UpdatePetWithForm)

	mapper.PutHandler(it.UploadFile)

}
