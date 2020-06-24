package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
	"net/http"
)

//
type ApiResponse struct {
	Code *int32 `json:"code,omitempty"`

	Type *string `json:"type,omitempty"`

	Message *string `json:"message,omitempty"`
}

// encode to json
func (it ApiResponse) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it ApiResponse) Json() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

type ApiResponseArray []ApiResponse

func (it *ApiResponse) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Code, it.Code == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Type, it.Type == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Message, it.Message == nil).
		Valid(factory) {
		return false
	}

	return true
}

func (it *ApiResponse) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *ApiResponseArray) Valid(factory swagger.ValidatorFactory) bool {
	for _, it := range *arr {
		if !factory.NewValidator(it.Code, it.Code == nil).
			Valid(factory) {
			return false
		}
		if !factory.NewValidator(it.Type, it.Type == nil).
			Valid(factory) {
			return false
		}
		if !factory.NewValidator(it.Message, it.Message == nil).
			Valid(factory) {
			return false
		}
	}
	return true
}

func (it *ApiResponseArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
