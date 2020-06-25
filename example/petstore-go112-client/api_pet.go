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
            Body   *Pet

        }

        /*
        Add a new pet to the store
        

          result: void
        */
        func (it *PetApi)AddPet(_client swagger.FetchClient, _request *PetApiAddPetRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.Body, _request.Body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Body' when calling AddPet")
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

            if _request.Body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.Body))
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
            PetId   *int64

            /*
            
            */
            ApiKey   *string

        }

        /*
        Deletes a pet
        

          result: void
        */
        func (it *PetApi)DeletePet(_client swagger.FetchClient, _request *PetApiDeletePetRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.PetId, _request.PetId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'PetId' when calling DeletePet")
            }
            if(!_client.NewValidator(_request.ApiKey, _request.ApiKey == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'ApiKey' when calling DeletePet")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.PetId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Delete"))
        }


            if _request.ApiKey != nil {
                _client.AddHeader("api_key", utils.ParameterToString(_request.ApiKey))
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
            Status   *[]string

        }

        /*
        Finds Pets by status
        Multiple status values can be provided with comma separated strings

          result: []Pet
        */
        func (it *PetApi)FindPetsByStatus(_client swagger.FetchClient, _request *PetApiFindPetsByStatusRequest, result *[]Pet) error {
            if(!_client.NewValidator(_request.Status, _request.Status == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Status' when calling FindPetsByStatus")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/findByStatus","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }

            if _request.Status != nil {
                _client.AddQueryParam("status", utils.ParameterToString(_request.Status))
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
            Tags   *[]string

        }

        /*
        Finds Pets by tags
        Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.

          result: []Pet
        */
        func (it *PetApi)FindPetsByTags(_client swagger.FetchClient, _request *PetApiFindPetsByTagsRequest, result *[]Pet) error {
            if(!_client.NewValidator(_request.Tags, _request.Tags == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Tags' when calling FindPetsByTags")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/findByTags","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }

            if _request.Tags != nil {
                _client.AddQueryParam("tags", utils.ParameterToString(_request.Tags))
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
            PetId   *int64

        }

        /*
        Find pet by ID
        Returns a single pet

          result: Pet
        */
        func (it *PetApi)GetPetById(_client swagger.FetchClient, _request *PetApiGetPetByIdRequest, result *Pet) error {
            if(!_client.NewValidator(_request.PetId, _request.PetId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'PetId' when calling GetPetById")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.PetId)), -1)
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
            Body   *Pet

        }

        /*
        Update an existing pet
        

          result: void
        */
        func (it *PetApi)UpdatePet(_client swagger.FetchClient, _request *PetApiUpdatePetRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.Body, _request.Body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Body' when calling UpdatePet")
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

            if _request.Body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.Body))
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
            PetId   *int64

            /*
            Updated name of the pet
            */
            Name   *string

            /*
            Updated status of the pet
            */
            Status   *string

        }

        /*
        Updates a pet in the store with form data
        

          result: void
        */
        func (it *PetApi)UpdatePetWithForm(_client swagger.FetchClient, _request *PetApiUpdatePetWithFormRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.PetId, _request.PetId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'PetId' when calling UpdatePetWithForm")
            }
            if(!_client.NewValidator(_request.Name, _request.Name == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Name' when calling UpdatePetWithForm")
            }
            if(!_client.NewValidator(_request.Status, _request.Status == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Status' when calling UpdatePetWithForm")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.PetId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Post"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        
            // form params
            formEnable = true
            if _request.Name != nil {
                localVarFormParams.Add("name", utils.ParameterToString(_request.Name))
            }
        
            // form params
            formEnable = true
            if _request.Status != nil {
                localVarFormParams.Add("status", utils.ParameterToString(_request.Status))
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
            PetId   *int64

            /*
            Additional data to pass to server
            */
            AdditionalMetadata   *string

            /*
            file to upload
            */
            File   *io.Reader

        }

        /*
        uploads an image
        

          result: ApiResponse
        */
        func (it *PetApi)UploadFile(_client swagger.FetchClient, _request *PetApiUploadFileRequest, result *ApiResponse) error {
            if(!_client.NewValidator(_request.PetId, _request.PetId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'PetId' when calling UploadFile")
            }
            if(!_client.NewValidator(_request.AdditionalMetadata, _request.AdditionalMetadata == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'AdditionalMetadata' when calling UploadFile")
            }
            if(!_client.NewValidator(_request.File, _request.File == nil).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'File' when calling UploadFile")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/pet/{petId}/uploadImage","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "petId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.PetId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Post"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        
            // form params
            formEnable = true
            if _request.AdditionalMetadata != nil {
                localVarFormParams.Add("additionalMetadata", utils.ParameterToString(_request.AdditionalMetadata))
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
