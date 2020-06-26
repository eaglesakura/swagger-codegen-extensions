package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

//
// User Interface for alternative definitions.
type IUser interface {
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

	// Remove Username property
	RemoveUsername()

	// Get Username property or default value.
	GetUsername() string

	// Get Username property or panic().
	RequireUsername() string

	// Set Username property.
	SetUsername(newUsername string)

	// Remove FirstName property
	RemoveFirstName()

	// Get FirstName property or default value.
	GetFirstName() string

	// Get FirstName property or panic().
	RequireFirstName() string

	// Set FirstName property.
	SetFirstName(newFirstName string)

	// Remove LastName property
	RemoveLastName()

	// Get LastName property or default value.
	GetLastName() string

	// Get LastName property or panic().
	RequireLastName() string

	// Set LastName property.
	SetLastName(newLastName string)

	// Remove Email property
	RemoveEmail()

	// Get Email property or default value.
	GetEmail() string

	// Get Email property or panic().
	RequireEmail() string

	// Set Email property.
	SetEmail(newEmail string)

	// Remove Password property
	RemovePassword()

	// Get Password property or default value.
	GetPassword() string

	// Get Password property or panic().
	RequirePassword() string

	// Set Password property.
	SetPassword(newPassword string)

	// Remove Phone property
	RemovePhone()

	// Get Phone property or default value.
	GetPhone() string

	// Get Phone property or panic().
	RequirePhone() string

	// Set Phone property.
	SetPhone(newPhone string)

	// Remove UserStatus property
	RemoveUserStatus()

	// Get UserStatus property or default value.
	GetUserStatus() int32

	// Get UserStatus property or panic().
	RequireUserStatus() int32

	// Set UserStatus property.
	SetUserStatus(newUserStatus int32)
}

//
type User struct {
	//
	Id *int64 `json:"id,omitempty"`
	//
	Username *string `json:"username,omitempty"`
	//
	FirstName *string `json:"firstName,omitempty"`
	//
	LastName *string `json:"lastName,omitempty"`
	//
	Email *string `json:"email,omitempty"`
	//
	Password *string `json:"password,omitempty"`
	//
	Phone *string `json:"phone,omitempty"`
	// User Status
	UserStatus *int32 `json:"userStatus,omitempty"`
}
type UserArray []User

// Update self from builder callback.
func (it *User) Update(builder func(ref *User)) *User {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *User) Copy(builder func(copied *User)) *User {
	result := *it
	builder(&result)
	return &result
}

// DeepCopy to result ptr.
func (it *User) MarshalCopy(result interface{}) error {
	body, _ := json.Marshal(it)
	err := json.Unmarshal(body, result)
	if err != nil {
		return xerrors.Errorf("User.MarshalCopy failed, to='%v': %w", result, err)
	}
	return nil
}

// Remove Id
func (it *User) RemoveId() {
	it.Id = nil
}

// Set Id
func (it *User) SetId(newId int64) {
	it.Id = &newId
}

// Require value of Id
func (it *User) RequireId() int64 {
	if it.Id == nil {
		panic(xerrors.Errorf("User.Id is nil"))
	}
	return *it.Id
}

// Get value of Id / or default
func (it *User) GetId() int64 {
	if it.Id != nil {
		return *it.Id
	}
	result := new(int64)
	return *result
}

// Remove Username
func (it *User) RemoveUsername() {
	it.Username = nil
}

// Set Username
func (it *User) SetUsername(newUsername string) {
	it.Username = &newUsername
}

// Require value of Username
func (it *User) RequireUsername() string {
	if it.Username == nil {
		panic(xerrors.Errorf("User.Username is nil"))
	}
	return *it.Username
}

// Get value of Username / or default
func (it *User) GetUsername() string {
	if it.Username != nil {
		return *it.Username
	}
	result := new(string)
	return *result
}

// Remove FirstName
func (it *User) RemoveFirstName() {
	it.FirstName = nil
}

// Set FirstName
func (it *User) SetFirstName(newFirstName string) {
	it.FirstName = &newFirstName
}

// Require value of FirstName
func (it *User) RequireFirstName() string {
	if it.FirstName == nil {
		panic(xerrors.Errorf("User.FirstName is nil"))
	}
	return *it.FirstName
}

// Get value of FirstName / or default
func (it *User) GetFirstName() string {
	if it.FirstName != nil {
		return *it.FirstName
	}
	result := new(string)
	return *result
}

// Remove LastName
func (it *User) RemoveLastName() {
	it.LastName = nil
}

// Set LastName
func (it *User) SetLastName(newLastName string) {
	it.LastName = &newLastName
}

// Require value of LastName
func (it *User) RequireLastName() string {
	if it.LastName == nil {
		panic(xerrors.Errorf("User.LastName is nil"))
	}
	return *it.LastName
}

// Get value of LastName / or default
func (it *User) GetLastName() string {
	if it.LastName != nil {
		return *it.LastName
	}
	result := new(string)
	return *result
}

// Remove Email
func (it *User) RemoveEmail() {
	it.Email = nil
}

// Set Email
func (it *User) SetEmail(newEmail string) {
	it.Email = &newEmail
}

// Require value of Email
func (it *User) RequireEmail() string {
	if it.Email == nil {
		panic(xerrors.Errorf("User.Email is nil"))
	}
	return *it.Email
}

// Get value of Email / or default
func (it *User) GetEmail() string {
	if it.Email != nil {
		return *it.Email
	}
	result := new(string)
	return *result
}

// Remove Password
func (it *User) RemovePassword() {
	it.Password = nil
}

// Set Password
func (it *User) SetPassword(newPassword string) {
	it.Password = &newPassword
}

// Require value of Password
func (it *User) RequirePassword() string {
	if it.Password == nil {
		panic(xerrors.Errorf("User.Password is nil"))
	}
	return *it.Password
}

// Get value of Password / or default
func (it *User) GetPassword() string {
	if it.Password != nil {
		return *it.Password
	}
	result := new(string)
	return *result
}

// Remove Phone
func (it *User) RemovePhone() {
	it.Phone = nil
}

// Set Phone
func (it *User) SetPhone(newPhone string) {
	it.Phone = &newPhone
}

// Require value of Phone
func (it *User) RequirePhone() string {
	if it.Phone == nil {
		panic(xerrors.Errorf("User.Phone is nil"))
	}
	return *it.Phone
}

// Get value of Phone / or default
func (it *User) GetPhone() string {
	if it.Phone != nil {
		return *it.Phone
	}
	result := new(string)
	return *result
}

// Remove UserStatus
func (it *User) RemoveUserStatus() {
	it.UserStatus = nil
}

// Set UserStatus
func (it *User) SetUserStatus(newUserStatus int32) {
	it.UserStatus = &newUserStatus
}

// Require value of UserStatus
func (it *User) RequireUserStatus() int32 {
	if it.UserStatus == nil {
		panic(xerrors.Errorf("User.UserStatus is nil"))
	}
	return *it.UserStatus
}

// Get value of UserStatus / or default
func (it *User) GetUserStatus() int32 {
	if it.UserStatus != nil {
		return *it.UserStatus
	}
	result := new(int32)
	return *result
}

// encode to json
func (it User) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it *User) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *User) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it *User) Valid() error {
	if it.Id != nil {
		if err := validationValue(it.Id); err != nil {
			return xerrors.Errorf("'User.Id' validation error: %w", err)
		}
	}
	if it.Username != nil {
		if err := validationValue(it.Username); err != nil {
			return xerrors.Errorf("'User.Username' validation error: %w", err)
		}
	}
	if it.FirstName != nil {
		if err := validationValue(it.FirstName); err != nil {
			return xerrors.Errorf("'User.FirstName' validation error: %w", err)
		}
	}
	if it.LastName != nil {
		if err := validationValue(it.LastName); err != nil {
			return xerrors.Errorf("'User.LastName' validation error: %w", err)
		}
	}
	if it.Email != nil {
		if err := validationValue(it.Email); err != nil {
			return xerrors.Errorf("'User.Email' validation error: %w", err)
		}
	}
	if it.Password != nil {
		if err := validationValue(it.Password); err != nil {
			return xerrors.Errorf("'User.Password' validation error: %w", err)
		}
	}
	if it.Phone != nil {
		if err := validationValue(it.Phone); err != nil {
			return xerrors.Errorf("'User.Phone' validation error: %w", err)
		}
	}
	if it.UserStatus != nil {
		if err := validationValue(it.UserStatus); err != nil {
			return xerrors.Errorf("'User.UserStatus' validation error: %w", err)
		}
	}

	return nil
}

func (it UserArray) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it UserArray) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *User) compilerDummy() {
	time.Now()
	_ = xerrors.Errorf("")

	var model User
	var iModel IUser = &model
	var swaggerModelRef swaggerModel = &model
	var swaggerResponseRef SwaggerResponse = &model
	var swaggerValidatableRef swaggerValidatable = &model

	var modelArray UserArray
	swaggerResponseRef = modelArray

	iModel = iModel
	swaggerModelRef = swaggerModelRef
	swaggerResponseRef = swaggerResponseRef
	swaggerValidatableRef = swaggerValidatableRef
}
