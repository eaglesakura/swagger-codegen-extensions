package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

type PetSTATUS string
type PetSTATUSArray []PetSTATUS

const PetSTATUS_available PetSTATUS = PetSTATUS("available")
const PetSTATUS_pending PetSTATUS = PetSTATUS("pending")
const PetSTATUS_sold PetSTATUS = PetSTATUS("sold")

var PetSTATUSPattern []string = []string{

	"available",
	"pending",
	"sold",
}

func (it PetSTATUS) Ptr() *PetSTATUS {
	return &it
}
func (it *PetSTATUS) Value() PetSTATUS {
	return *it
}
func (it *PetSTATUS) Valid() bool {
	if it == nil {
		return false
	}
	value := string(*it)
	for _, v := range PetSTATUSPattern {
		if v == value {
			return true
		}
	}
	return false
}

//
type Pet struct {
	Id *int64 `json:"id,omitempty"`

	Category *Category `json:"category,omitempty"`

	Name *string `json:"name"`

	PhotoUrls *[]string `json:"photoUrls"`

	Tags *[]Tag `json:"tags,omitempty"`

	// pet status in the store
	Status *PetSTATUS `json:"status,omitempty"`
}
type PetArray []Pet

// Update self from builder callback.
func (it *Pet) Update(builder func(ref *Pet)) *Pet {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *Pet) Copy(builder func(copied *Pet)) *Pet {
	result := *it
	builder(&result)
	return &result
}

// DeepCopy to result ptr.
func (it *Pet) MarshalCopy(result interface{}) error {
	body, _ := json.Marshal(it)
	err := json.Unmarshal(body, result)
	if err != nil {
		return xerrors.Errorf("Pet.MarshalCopy failed, to='%v', %w", result, err)
	}
	return nil
}

// Set Id
func (it *Pet) SetId(newId int64) {
	it.Id = &newId
}

// Require value of Id
func (it *Pet) RequireId() int64 {
	if it.Id == nil {
		panic(xerrors.Errorf("Pet.Id is nil"))
	}
	return *it.Id
}

// Get value of Id / or default
func (it *Pet) GetId() int64 {
	if it.Id != nil {
		return *it.Id
	}
	result := new(int64)
	return *result
}

// Set Category
func (it *Pet) SetCategory(newCategory Category) {
	it.Category = &newCategory
}

// Require value of Category
func (it *Pet) RequireCategory() Category {
	if it.Category == nil {
		panic(xerrors.Errorf("Pet.Category is nil"))
	}
	return *it.Category
}

// Get value of Category / or default
func (it *Pet) GetCategory() Category {
	if it.Category != nil {
		return *it.Category
	}
	result := new(Category)
	return *result
}

// Set Name
func (it *Pet) SetName(newName string) {
	it.Name = &newName
}

// Require value of Name
func (it *Pet) RequireName() string {
	if it.Name == nil {
		panic(xerrors.Errorf("Pet.Name is nil"))
	}
	return *it.Name
}

// Get value of Name / or default
func (it *Pet) GetName() string {
	if it.Name != nil {
		return *it.Name
	}
	result := new(string)
	return *result
}

// Set PhotoUrls
func (it *Pet) SetPhotoUrls(newPhotoUrls []string) {
	it.PhotoUrls = &newPhotoUrls
}

// Require value of PhotoUrls
func (it *Pet) RequirePhotoUrls() []string {
	if it.PhotoUrls == nil {
		panic(xerrors.Errorf("Pet.PhotoUrls is nil"))
	}
	return *it.PhotoUrls
}

// Get value of PhotoUrls / or default
func (it *Pet) GetPhotoUrls() []string {
	if it.PhotoUrls != nil {
		return *it.PhotoUrls
	}
	result := new([]string)
	return *result
}

// Set Tags
func (it *Pet) SetTags(newTags []Tag) {
	it.Tags = &newTags
}

// Require value of Tags
func (it *Pet) RequireTags() []Tag {
	if it.Tags == nil {
		panic(xerrors.Errorf("Pet.Tags is nil"))
	}
	return *it.Tags
}

// Get value of Tags / or default
func (it *Pet) GetTags() []Tag {
	if it.Tags != nil {
		return *it.Tags
	}
	result := new([]Tag)
	return *result
}

// Set Status
func (it *Pet) SetStatus(newStatus PetSTATUS) {
	it.Status = &newStatus
}

// Require value of Status
func (it *Pet) RequireStatus() PetSTATUS {
	if it.Status == nil {
		panic(xerrors.Errorf("Pet.Status is nil"))
	}
	return *it.Status
}

// Get value of Status / or default
func (it *Pet) GetStatus() PetSTATUS {
	if it.Status != nil {
		return *it.Status
	}
	result := new(PetSTATUS)
	return *result
}

// encode to json
func (it Pet) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it *Pet) Json() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

func (it *Pet) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it *Pet) Valid() error {
	if it.Name == nil {
		return xerrors.Errorf("'Pet.Name' is required")
	}
	if it.PhotoUrls == nil {
		return xerrors.Errorf("'Pet.PhotoUrls' is required")
	}
	if it.Status != nil && !it.Status.Valid() {
		return xerrors.Errorf("'Pet.Status' is invalid")
	}

	return nil
}

func (it *Pet) this_is_call_dummy() {
	time.Now()
	xerrors.Errorf("")

	var model Pet
	var swaggerModelRef swaggerModel = &model
	var swaggerResponseRef SwaggerResponse = &model
	swaggerModelRef = swaggerModelRef
	swaggerResponseRef = swaggerResponseRef
}
