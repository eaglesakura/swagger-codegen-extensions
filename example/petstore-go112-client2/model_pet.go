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

const PetSTATUS_available = PetSTATUS("available")
const PetSTATUS_pending = PetSTATUS("pending")
const PetSTATUS_sold = PetSTATUS("sold")

func (it PetSTATUS) String() string {
	return string(it)
}

// Pattern for PetSTATUS
var PetSTATUSValues = []PetSTATUS{

	PetSTATUS_available,

	PetSTATUS_pending,

	PetSTATUS_sold,
}

func (it *PetSTATUS) Valid() error {
	if it == nil {
		return xerrors.Errorf("PetSTATUS(nil)")
	}
	switch *it {

	case PetSTATUS_available:
		return nil

	case PetSTATUS_pending:
		return nil

	case PetSTATUS_sold:
		return nil

	}
	return xerrors.Errorf("invalid PetSTATUS(%v)", string(*it))
}

//
// Pet Interface for alternative definitions.
type IPet interface {
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

	// Remove Category property
	RemoveCategory()

	// Get Category property or default value.
	GetCategory() Category

	// Get Category property or panic().
	RequireCategory() Category

	// Set Category property.
	SetCategory(newCategory Category)

	// Remove Name property
	RemoveName()

	// Get Name property or default value.
	GetName() string

	// Get Name property or panic().
	RequireName() string

	// Set Name property.
	SetName(newName string)

	// Remove PhotoUrls property
	RemovePhotoUrls()

	// Get PhotoUrls property or default value.
	GetPhotoUrls() []string

	// Get PhotoUrls property or panic().
	RequirePhotoUrls() []string

	// Set PhotoUrls property.
	SetPhotoUrls(newPhotoUrls []string)

	// Remove Tags property
	RemoveTags()

	// Get Tags property or default value.
	GetTags() []Tag

	// Get Tags property or panic().
	RequireTags() []Tag

	// Set Tags property.
	SetTags(newTags []Tag)

	// Remove Status property
	RemoveStatus()

	// Get Status property or default value.
	GetStatus() PetSTATUS

	// Get Status property or panic().
	RequireStatus() PetSTATUS

	// Set Status property.
	SetStatus(newStatus PetSTATUS)
}

//
type Pet struct {
	//
	Id *int64 `json:"id,omitempty"`
	//
	Category *Category `json:"category,omitempty"`
	//
	Name *string `json:"name"`
	//
	PhotoUrls *[]string `json:"photoUrls"`
	//
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

// Remove Id
func (it *Pet) RemoveId() {
	it.Id = nil
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

// Remove Category
func (it *Pet) RemoveCategory() {
	it.Category = nil
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

// Remove Name
func (it *Pet) RemoveName() {
	it.Name = nil
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

// Remove PhotoUrls
func (it *Pet) RemovePhotoUrls() {
	it.PhotoUrls = nil
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

// Remove Tags
func (it *Pet) RemoveTags() {
	it.Tags = nil
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

// Remove Status
func (it *Pet) RemoveStatus() {
	it.Status = nil
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
func (it *Pet) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *Pet) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it *Pet) Valid() error {
	if it.Id != nil {
		if err := validationValue(it.Id); err != nil {
			return xerrors.Errorf("'Pet.Id' validation error, %w", err)
		}
	}
	if it.Category != nil {
		if err := validationValue(it.Category); err != nil {
			return xerrors.Errorf("'Pet.Category' validation error, %w", err)
		}
	}
	if it.Name == nil {
		return xerrors.Errorf("'Pet.Name' is required")
	}
	if it.Name != nil {
		if err := validationValue(it.Name); err != nil {
			return xerrors.Errorf("'Pet.Name' validation error, %w", err)
		}
	}
	if it.PhotoUrls == nil {
		return xerrors.Errorf("'Pet.PhotoUrls' is required")
	}
	if it.PhotoUrls != nil {
		if err := validationValue(it.PhotoUrls); err != nil {
			return xerrors.Errorf("'Pet.PhotoUrls' validation error, %w", err)
		}
	}
	if it.Tags != nil {
		if err := validationValue(it.Tags); err != nil {
			return xerrors.Errorf("'Pet.Tags' validation error, %w", err)
		}
	}
	if it.Status != nil {
		if err := validationValue(it.Status); err != nil {
			return xerrors.Errorf("'Pet.Status' validation error, %w", err)
		}
	}

	return nil
}

func (it PetArray) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it PetArray) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *Pet) compilerDummy() {
	time.Now()
	_ = xerrors.Errorf("")

	var model Pet
	var iModel IPet = &model
	var swaggerModelRef swaggerModel = &model
	var swaggerResponseRef SwaggerResponse = &model
	var swaggerValidatableRef swaggerValidatable = &model

	var modelArray PetArray
	swaggerResponseRef = modelArray

	iModel = iModel
	swaggerModelRef = swaggerModelRef
	swaggerResponseRef = swaggerResponseRef
	swaggerValidatableRef = swaggerValidatableRef
}
