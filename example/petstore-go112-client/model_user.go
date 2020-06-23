package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions



import (
    "time"
    "net/http"
    "encoding/json"
    "github.com/eaglesakura/swagger-go-core"
)



// 
type User struct {

	Id *int64 `json:"id,omitempty"`

	Username *string `json:"username,omitempty"`

	FirstName *string `json:"firstName,omitempty"`

	LastName *string `json:"lastName,omitempty"`

	Email *string `json:"email,omitempty"`

	Password *string `json:"password,omitempty"`

	Phone *string `json:"phone,omitempty"`

	// User Status
	UserStatus *int32 `json:"userStatus,omitempty"`
}

// encode to json
func (it User)String() string {
    buf, _ := json.Marshal(it)
    return string(buf)
}
func (it User)Json() string {
    buf, _ := json.Marshal(it)
    return string(buf)
}

type UserArray []User

func (it *User) Valid(factory swagger.ValidatorFactory) bool {
    if(!factory.NewValidator(it.Id, it.Id == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Username, it.Username == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.FirstName, it.FirstName == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.LastName, it.LastName == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Email, it.Email == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Password, it.Password == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Phone, it.Phone == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.UserStatus, it.UserStatus == nil).
                
                
                
                Valid(factory)) {
        return false
    }

    return true
}

func (it *User) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil  {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *UserArray) Valid(factory swagger.ValidatorFactory) bool {
    for _, it := range *arr {
        if(!factory.NewValidator(it.Id, it.Id == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Username, it.Username == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.FirstName, it.FirstName == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.LastName, it.LastName == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Email, it.Email == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Password, it.Password == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Phone, it.Phone == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.UserStatus, it.UserStatus == nil).
        
        
        
        Valid(factory)) {
        return false
        }
    }
    return true
}

func (it *UserArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
    writer.WriteHeader(200)
    if err := producer.Produce(writer, it); err != nil  {
        panic(err) // let the recovery middleware deal with this
    }
}


func (it *User)this_is_call_dummy() {
    time.Now()
}




