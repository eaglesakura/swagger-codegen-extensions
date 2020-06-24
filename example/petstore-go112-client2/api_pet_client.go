package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"net/http"
)

// petstore
type PetApi struct {
	// API Endpoint
	// e.g.) "https://example.com/"
	Endpoint string

	// http.Client prefetch intercept function.
	Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

// Update self from builder callback.
func (it *PetApi) Update(builder func(ref *PetApi)) *PetApi {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *PetApi) Copy(builder func(copied *PetApi)) *PetApi {
	result := *it
	builder(&result)
	return &result
}

func NewPetApi() *PetApi {
	return &PetApi{
		Endpoint: "",
	}
}

func (it *PetApi) AddPetPost(builder func(*PetApiAddPetPostRequest)) *PetApiAddPetPostRequest {
	result := &PetApiAddPetPostRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *PetApi) DeletePetDelete(builder func(*PetApiDeletePetDeleteRequest)) *PetApiDeletePetDeleteRequest {
	result := &PetApiDeletePetDeleteRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *PetApi) FindPetsByStatusGet(builder func(*PetApiFindPetsByStatusGetRequest)) *PetApiFindPetsByStatusGetRequest {
	result := &PetApiFindPetsByStatusGetRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *PetApi) FindPetsByTagsGet(builder func(*PetApiFindPetsByTagsGetRequest)) *PetApiFindPetsByTagsGetRequest {
	result := &PetApiFindPetsByTagsGetRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *PetApi) GetPetByIdGet(builder func(*PetApiGetPetByIdGetRequest)) *PetApiGetPetByIdGetRequest {
	result := &PetApiGetPetByIdGetRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *PetApi) UpdatePetPut(builder func(*PetApiUpdatePetPutRequest)) *PetApiUpdatePetPutRequest {
	result := &PetApiUpdatePetPutRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *PetApi) UpdatePetWithFormPost(builder func(*PetApiUpdatePetWithFormPostRequest)) *PetApiUpdatePetWithFormPostRequest {
	result := &PetApiUpdatePetWithFormPostRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
func (it *PetApi) UploadFilePost(builder func(*PetApiUploadFilePostRequest)) *PetApiUploadFilePostRequest {
	result := &PetApiUploadFilePostRequest{
		Endpoint:  it.Endpoint,
		Intercept: it.Intercept,
	}
	builder(result)
	return result
}
