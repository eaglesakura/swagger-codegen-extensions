package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
    "net/http"
    "time"
)






// 
// Tag Interface for alternative definitions.
type ITag interface {
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
type Tag struct {
        // 
        Id *int64 `json:"id,omitempty"`
        // 
        Name *string `json:"name,omitempty"`
}
type TagArray []Tag

// Update self from builder callback.
func (it *Tag)Update(builder func(ref *Tag)) *Tag {
    builder(it)
    return it
}

// Copy and Update self from builder callback.
func (it *Tag)Copy(builder func(copied *Tag)) *Tag {
    result := *it
    builder(&result)
    return &result
}

// DeepCopy to result ptr.
func (it *Tag)MarshalCopy(result interface{}) error {
    body, _ := json.Marshal(it)
    err := json.Unmarshal(body, result)
    if err != nil {
        return xerrors.Errorf("Tag.MarshalCopy failed, to='%v': %w", result, err)
    }
    return nil
}

// Remove Id
func (it *Tag)RemoveId() {
    it.Id = nil
}

// Set Id
func (it *Tag)SetId(newId int64) {
    it.Id = &newId
}

// Require value of Id
func (it *Tag)RequireId() int64 {
    if it.Id == nil {
        panic(xerrors.Errorf("Tag.Id is nil"))
    }
    return *it.Id
}

// Get value of Id / or default
func (it *Tag)GetId() int64 {
    if it.Id != nil {
        return *it.Id
    }
    result := new(int64)
    return *result
}
// Remove Name
func (it *Tag)RemoveName() {
    it.Name = nil
}

// Set Name
func (it *Tag)SetName(newName string) {
    it.Name = &newName
}

// Require value of Name
func (it *Tag)RequireName() string {
    if it.Name == nil {
        panic(xerrors.Errorf("Tag.Name is nil"))
    }
    return *it.Name
}

// Get value of Name / or default
func (it *Tag)GetName() string {
    if it.Name != nil {
        return *it.Name
    }
    result := new(string)
    return *result
}

// encode to json
func (it Tag)String() string {
    buf, _ := json.Marshal(it)
    return string(buf)
}
func (it *Tag)Json() []byte {
    buf, _ := json.Marshal(it)
    return buf
}

func (it *Tag) Write(writer http.ResponseWriter, request *http.Request) {
    buf, _ := json.Marshal(it)
    _, _ = writer.Write(buf)
}

func (it *Tag) Valid() error {
        if it.Id != nil {
            if err := validationValue(it.Id); err != nil {
                return xerrors.Errorf("'Tag.Id' validation error: %w", err)
            }
        }
        if it.Name != nil {
            if err := validationValue(it.Name); err != nil {
                return xerrors.Errorf("'Tag.Name' validation error: %w", err)
            }
        }

    return nil
}

func (it TagArray) Write(writer http.ResponseWriter, request *http.Request) {
    buf, _ := json.Marshal(it)
    _, _ = writer.Write(buf)
}

func (it TagArray) Json() []byte {
    buf, _ := json.Marshal(it)
    return buf
}


func (it *Tag)compilerDummy() {
    time.Now()
    _ = xerrors.Errorf("")

    var model Tag
    var iModel ITag = &model
    var swaggerModelRef swaggerModel = &model
    var swaggerResponseRef SwaggerResponse = &model
    var swaggerValidatableRef swaggerValidatable = &model

    var modelArray TagArray
    swaggerResponseRef = modelArray

    iModel = iModel
    swaggerModelRef = swaggerModelRef
    swaggerResponseRef = swaggerResponseRef
    swaggerValidatableRef = swaggerValidatableRef
}




