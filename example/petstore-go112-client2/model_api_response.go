package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"time"
)

//
type ApiResponse struct {
	Code *int32 `json:"code,omitempty"`

	Type *string `json:"type,omitempty"`

	Message *string `json:"message,omitempty"`
}

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
func (it ApiResponse) Json() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

type ApiResponseArray []ApiResponse

func (it *ApiResponse) Valid() error {

	return nil
}

func (it *ApiResponse) this_is_call_dummy() {
	time.Now()
	xerrors.Errorf("")
}
