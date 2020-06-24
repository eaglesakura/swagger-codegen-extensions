package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"time"
)

//
type Tag struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

// Update self from builder callback.
func (it *Tag) Update(builder func(ref *Tag)) *Tag {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *Tag) Copy(builder func(copied *Tag)) *Tag {
	result := *it
	builder(&result)
	return &result
}

// Set Id
func (it *Tag) SetId(newId int64) {
	it.Id = &newId
}

// Require value of Id
func (it *Tag) RequireId() int64 {
	if it.Id == nil {
		panic(xerrors.Errorf("Tag.Id is nil"))
	}
	return *it.Id
}

// Get value of Id / or default
func (it *Tag) GetId() int64 {
	if it.Id != nil {
		return *it.Id
	}
	result := new(int64)
	return *result
}

// Set Name
func (it *Tag) SetName(newName string) {
	it.Name = &newName
}

// Require value of Name
func (it *Tag) RequireName() string {
	if it.Name == nil {
		panic(xerrors.Errorf("Tag.Name is nil"))
	}
	return *it.Name
}

// Get value of Name / or default
func (it *Tag) GetName() string {
	if it.Name != nil {
		return *it.Name
	}
	result := new(string)
	return *result
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

func (it *Tag) Valid() error {

	return nil
}

func (it *Tag) this_is_call_dummy() {
	time.Now()
	xerrors.Errorf("")
}
