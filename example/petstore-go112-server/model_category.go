package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
	"net/http"
)

//
type Category struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

// encode to json
func (it Category) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it Category) Json() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

type CategoryArray []Category

func (it *Category) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Id, it.Id == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Name, it.Name == nil).
		Valid(factory) {
		return false
	}

	return true
}

func (it *Category) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *CategoryArray) Valid(factory swagger.ValidatorFactory) bool {
	for _, it := range *arr {
		if !factory.NewValidator(it.Id, it.Id == nil).
			Valid(factory) {
			return false
		}
		if !factory.NewValidator(it.Name, it.Name == nil).
			Valid(factory) {
			return false
		}
	}
	return true
}

func (it *CategoryArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
