package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
    "fmt"
    "io"
    "net/url"
    "strings"
    "github.com/eaglesakura/swagger-go-core"
    "github.com/eaglesakura/swagger-go-core/errors"
    "github.com/eaglesakura/swagger-go-core/utils"
)


const PetApi_BasePath string = "/v2"
type PetApi struct {
    BasePath string
}


func NewPetApi() *PetApi {
    return &PetApi{
        BasePath:PetApi_BasePath,
    }
}

        /*
        Add a new pet to the store
        
        */
        type PetApiAddPetRequest struct {
            /*
            Pet object that needs to be added to the store
            */
            body   *Pet

        }

        /*
        Add a new pet to the store
        

          result: void
        */
        func (it *PetApi)AddPet(_client swagger.FetchClient, _request *PetApiAddPetRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.body, _request.body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'body' when calling AddPet")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Post"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }

            if _request.body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.body))
            }

            return _client.Fetch(result)
        }
        /*
        Deletes a pet
        
        */
        type PetApiDeletePetRequest struct {
            /*
            Pet id to delete
            */
            petId   *int64

            /*
            
            */
            apiKey   *string

        }

        /*
        Deletes a pet
        

          result: void
        */
        func (it *PetApi)DeletePet(_client swagger.FetchClient, _request *PetApiDeletePetRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.petId, _request.petId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'petId' when calling DeletePet")
            }
            if(!_client.NewValidator(_request.apiKey, _request.apiKey == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'apiKey' when calling DeletePet")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.petId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Delete"))
        }


            if _request.apiKey != nil {
                _client.AddHeader("api_key", utils.ParameterToString(_request.apiKey))
            }

        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Finds Pets by status
        Multiple status values can be provided with comma separated strings
        */
        type PetApiFindPetsByStatusRequest struct {
            /*
            Status values that need to be considered for filter
            */
            status   *[]string

        }

        /*
        Finds Pets by status
        Multiple status values can be provided with comma separated strings

          result: []Pet
        */
        func (it *PetApi)FindPetsByStatus(_client swagger.FetchClient, _request *PetApiFindPetsByStatusRequest, result *[]Pet) error {
            if(!_client.NewValidator(_request.status, _request.status == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'status' when calling FindPetsByStatus")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/findByStatus","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }

            if _request.status != nil {
                _client.AddQueryParam("status", utils.ParameterToString(_request.status))
            }


        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Finds Pets by tags
        Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
        */
        type PetApiFindPetsByTagsRequest struct {
            /*
            Tags to filter by
            */
            tags   *[]string

        }

        /*
        Finds Pets by tags
        Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.

          result: []Pet
        */
        func (it *PetApi)FindPetsByTags(_client swagger.FetchClient, _request *PetApiFindPetsByTagsRequest, result *[]Pet) error {
            if(!_client.NewValidator(_request.tags, _request.tags == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'tags' when calling FindPetsByTags")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/findByTags","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }

            if _request.tags != nil {
                _client.AddQueryParam("tags", utils.ParameterToString(_request.tags))
            }


        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Find pet by ID
        Returns a single pet
        */
        type PetApiGetPetByIdRequest struct {
            /*
            ID of pet to return
            */
            petId   *int64

        }

        /*
        Find pet by ID
        Returns a single pet

          result: Pet
        */
        func (it *PetApi)GetPetById(_client swagger.FetchClient, _request *PetApiGetPetByIdRequest, result *Pet) error {
            if(!_client.NewValidator(_request.petId, _request.petId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'petId' when calling GetPetById")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.petId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Update an existing pet
        
        */
        type PetApiUpdatePetRequest struct {
            /*
            Pet object that needs to be added to the store
            */
            body   *Pet

        }

        /*
        Update an existing pet
        

          result: void
        */
        func (it *PetApi)UpdatePet(_client swagger.FetchClient, _request *PetApiUpdatePetRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.body, _request.body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'body' when calling UpdatePet")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Put"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }

            if _request.body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.body))
            }

            return _client.Fetch(result)
        }
        /*
        Updates a pet in the store with form data
        
        */
        type PetApiUpdatePetWithFormRequest struct {
            /*
            ID of pet that needs to be updated
            */
            petId   *int64

            /*
            Updated name of the pet
            */
            name   *string

            /*
            Updated status of the pet
            */
            status   *string

        }

        /*
        Updates a pet in the store with form data
        

          result: void
        */
        func (it *PetApi)UpdatePetWithForm(_client swagger.FetchClient, _request *PetApiUpdatePetWithFormRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.petId, _request.petId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'petId' when calling UpdatePetWithForm")
            }
            if(!_client.NewValidator(_request.name, _request.name == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'name' when calling UpdatePetWithForm")
            }
            if(!_client.NewValidator(_request.status, _request.status == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'status' when calling UpdatePetWithForm")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.petId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Post"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        
            // form params
            formEnable = true
            if _request.name != nil {
                localVarFormParams.Add("name", utils.ParameterToString(_request.name))
            }
        
            // form params
            formEnable = true
            if _request.status != nil {
                localVarFormParams.Add("status", utils.ParameterToString(_request.status))
            }
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        uploads an image
        
        */
        type PetApiUploadFileRequest struct {
            /*
            ID of pet to update
            */
            petId   *int64

            /*
            Additional data to pass to server
            */
            additionalMetadata   *string

            /*
            file to upload
            */
            file   *io.Reader

        }

        /*
        uploads an image
        

          result: ApiResponse
        */
        func (it *PetApi)UploadFile(_client swagger.FetchClient, _request *PetApiUploadFileRequest, result *ApiResponse) error {
            if(!_client.NewValidator(_request.petId, _request.petId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'petId' when calling UploadFile")
            }
            if(!_client.NewValidator(_request.additionalMetadata, _request.additionalMetadata == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'additionalMetadata' when calling UploadFile")
            }
            if(!_client.NewValidator(_request.file, _request.file == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'file' when calling UploadFile")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}/uploadImage","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.petId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Post"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        
            // form params
            formEnable = true
            if _request.additionalMetadata != nil {
                localVarFormParams.Add("additionalMetadata", utils.ParameterToString(_request.additionalMetadata))
            }
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }


func (it *PetApi)this_is_call_dummy() {
    v := url.Values{}
    v.Add("Key", "Value")

    errors.New(0, "stub")
    strings.ToUpper("")
    fmt.Sprintf("%v", "")
    io.ReadAtLeast(nil, nil, 0)
}
