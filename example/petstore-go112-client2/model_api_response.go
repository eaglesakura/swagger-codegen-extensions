package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

//
// ApiResponse Interface for alternative definitions.
type IApiResponse interface {
	// JsonCopy to result ptr.
	MarshalCopy(result interface{}) error

	// Marshal json string.
	Json() []byte

	// Remove Code property
	RemoveCode()

	// Get Code property or default value.
	GetCode() int32

	// Get Code property or panic().
	RequireCode() int32

	// Set Code property.
	SetCode(newCode int32)

	// Remove Type property
	RemoveType()

	// Get Type property or default value.
	GetType() string

	// Get Type property or panic().
	RequireType() string

	// Set Type property.
	SetType(newType string)

	// Remove Message property
	RemoveMessage()

	// Get Message property or default value.
	GetMessage() string

	// Get Message property or panic().
	RequireMessage() string

	// Set Message property.
	SetMessage(newMessage string)
}

//
type ApiResponse struct {
	//
	Code *int32 `json:"code,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	//
	Message *string `json:"message,omitempty"`
}
type ApiResponseArray []ApiResponse

// Update self from builder callback.
func (it *ApiResponse) Update(builder func(ref *ApiResponse)) *ApiResponse {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *ApiResponse) Copy(builder func(copied *ApiResponse)) *ApiResponse {
	result := *it
	builder(&result)
	return &result
}

// DeepCopy to result ptr.
func (it *ApiResponse) MarshalCopy(result interface{}) error {
	body, _ := json.Marshal(it)
	err := json.Unmarshal(body, result)
	if err != nil {
		return xerrors.Errorf("ApiResponse.MarshalCopy failed, to='%v', %w", result, err)
	}
	return nil
}

// Remove Code
func (it *ApiResponse) RemoveCode() {
	it.Code = nil
}

// Set Code
func (it *ApiResponse) SetCode(newCode int32) {
	it.Code = &newCode
}

// Require value of Code
func (it *ApiResponse) RequireCode() int32 {
	if it.Code == nil {
		panic(xerrors.Errorf("ApiResponse.Code is nil"))
	}
	return *it.Code
}

// Get value of Code / or default
func (it *ApiResponse) GetCode() int32 {
	if it.Code != nil {
		return *it.Code
	}
	result := new(int32)
	return *result
}

// Remove Type
func (it *ApiResponse) RemoveType() {
	it.Type = nil
}

// Set Type
func (it *ApiResponse) SetType(newType string) {
	it.Type = &newType
}

// Require value of Type
func (it *ApiResponse) RequireType() string {
	if it.Type == nil {
		panic(xerrors.Errorf("ApiResponse.Type is nil"))
	}
	return *it.Type
}

// Get value of Type / or default
func (it *ApiResponse) GetType() string {
	if it.Type != nil {
		return *it.Type
	}
	result := new(string)
	return *result
}

// Remove Message
func (it *ApiResponse) RemoveMessage() {
	it.Message = nil
}

// Set Message
func (it *ApiResponse) SetMessage(newMessage string) {
	it.Message = &newMessage
}

// Require value of Message
func (it *ApiResponse) RequireMessage() string {
	if it.Message == nil {
		panic(xerrors.Errorf("ApiResponse.Message is nil"))
	}
	return *it.Message
}

// Get value of Message / or default
func (it *ApiResponse) GetMessage() string {
	if it.Message != nil {
		return *it.Message
	}
	result := new(string)
	return *result
}

// encode to json
func (it ApiResponse) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it *ApiResponse) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *ApiResponse) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it *ApiResponse) Valid() error {
	if it.Code != nil {
		if err := validationValue(it.Code); err != nil {
			return xerrors.Errorf("'ApiResponse.Code' validation error, %w", err)
		}
	}
	if it.Type != nil {
		if err := validationValue(it.Type); err != nil {
			return xerrors.Errorf("'ApiResponse.Type' validation error, %w", err)
		}
	}
	if it.Message != nil {
		if err := validationValue(it.Message); err != nil {
			return xerrors.Errorf("'ApiResponse.Message' validation error, %w", err)
		}
	}

	return nil
}

func (it ApiResponseArray) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it ApiResponseArray) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *ApiResponse) compilerDummy() {
	time.Now()
	_ = xerrors.Errorf("")

	var model ApiResponse
	var iModel IApiResponse = &model
	var swaggerModelRef swaggerModel = &model
	var swaggerResponseRef SwaggerResponse = &model
	var swaggerValidatableRef swaggerValidatable = &model

	var modelArray ApiResponseArray
	swaggerResponseRef = modelArray

	iModel = iModel
	swaggerModelRef = swaggerModelRef
	swaggerResponseRef = swaggerResponseRef
	swaggerValidatableRef = swaggerValidatableRef
}
