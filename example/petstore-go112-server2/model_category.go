package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

//
// Category Interface for alternative definitions.
type ICategory interface {
	// JsonCopy to result ptr.
	MarshalCopy(result interface{}) error

	// Marshal json string.
	Json() []byte

	// Remove Id property
	RemoveId()

	// Get Id property or default value.
	GetId() int64

	// Get Id property or panic().
	RequireId() int64

	// Set Id property.
	SetId(newId int64)

	// Remove Name property
	RemoveName()

	// Get Name property or default value.
	GetName() string

	// Get Name property or panic().
	RequireName() string

	// Set Name property.
	SetName(newName string)
}

//
type Category struct {
	//
	Id *int64 `json:"id,omitempty"`
	//
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

// Remove Id
func (it *Category) RemoveId() {
	it.Id = nil
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

// Remove Name
func (it *Category) RemoveName() {
	it.Name = nil
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
func (it *Category) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *Category) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it *Category) Valid() error {
	if it.Id != nil {
		if err := validationValue(it.Id); err != nil {
			return xerrors.Errorf("'Category.Id' validation error, %w", err)
		}
	}
	if it.Name != nil {
		if err := validationValue(it.Name); err != nil {
			return xerrors.Errorf("'Category.Name' validation error, %w", err)
		}
	}

	return nil
}

func (it CategoryArray) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it CategoryArray) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *Category) compilerDummy() {
	time.Now()
	_ = xerrors.Errorf("")

	var model Category
	var iModel ICategory = &model
	var swaggerModelRef swaggerModel = &model
	var swaggerResponseRef SwaggerResponse = &model
	var swaggerValidatableRef swaggerValidatable = &model

	var modelArray CategoryArray
	swaggerResponseRef = modelArray

	iModel = iModel
	swaggerModelRef = swaggerModelRef
	swaggerResponseRef = swaggerResponseRef
	swaggerValidatableRef = swaggerValidatableRef
}
