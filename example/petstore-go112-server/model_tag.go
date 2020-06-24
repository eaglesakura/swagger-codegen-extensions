package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
	"net/http"
)

//
type Tag struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

// encode to json
func (it Tag) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it Tag) Json() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

type TagArray []Tag

func (it *Tag) Valid(factory swagger.ValidatorFactory) bool {
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

func (it *Tag) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *TagArray) Valid(factory swagger.ValidatorFactory) bool {
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

func (it *TagArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
