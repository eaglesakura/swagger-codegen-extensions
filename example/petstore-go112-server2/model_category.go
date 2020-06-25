package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

//
type Category struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}
type CategoryArray []Category

// Update self from builder callback.
func (it *Category) Update(builder func(ref *Category)) *Category {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *Category) Copy(builder func(copied *Category)) *Category {
	result := *it
	builder(&result)
	return &result
}

// DeepCopy to result ptr.
func (it *Category) MarshalCopy(result interface{}) error {
	body, _ := json.Marshal(it)
	err := json.Unmarshal(body, result)
	if err != nil {
		return xerrors.Errorf("Category.MarshalCopy failed, to='%v', %w", result, err)
	}
	return nil
}

// Set Id
func (it *Category) SetId(newId int64) {
	it.Id = &newId
}

// Require value of Id
func (it *Category) RequireId() int64 {
	if it.Id == nil {
		panic(xerrors.Errorf("Category.Id is nil"))
	}
	return *it.Id
}

// Get value of Id / or default
func (it *Category) GetId() int64 {
	if it.Id != nil {
		return *it.Id
	}
	result := new(int64)
	return *result
}

// Set Name
func (it *Category) SetName(newName string) {
	it.Name = &newName
}

// Require value of Name
func (it *Category) RequireName() string {
	if it.Name == nil {
		panic(xerrors.Errorf("Category.Name is nil"))
	}
	return *it.Name
}

// Get value of Name / or default
func (it *Category) GetName() string {
	if it.Name != nil {
		return *it.Name
	}
	result := new(string)
	return *result
}

// encode to json
func (it Category) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it *Category) Json() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

func (it *Category) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it *Category) Valid() error {

	return nil
}

func (it *Category) this_is_call_dummy() {
	time.Now()
	xerrors.Errorf("")

	var model Category
	var swaggerModelRef swaggerModel = &model
	var swaggerResponseRef SwaggerResponse = &model
	swaggerModelRef = swaggerModelRef
	swaggerResponseRef = swaggerResponseRef
}
