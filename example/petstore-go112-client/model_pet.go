package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions



import (
    "time"
    "net/http"
    "encoding/json"
    "github.com/eaglesakura/swagger-go-core"
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
func (it PetSTATUS)Ptr() *PetSTATUS {
    return &it
}
func (it *PetSTATUS)Value() PetSTATUS {
    return *it
}
func (it *PetSTATUS)Valid(pattern []string) bool {
    if it == nil {
        return false
    }
    value := string(*it)
    for _, v := range pattern {
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

// encode to json
func (it Pet)String() string {
    buf, _ := json.Marshal(it)
    return string(buf)
}
func (it Pet)Json() string {
    buf, _ := json.Marshal(it)
    return string(buf)
}

type PetArray []Pet

func (it *Pet) Valid(factory swagger.ValidatorFactory) bool {
    if(!factory.NewValidator(it.Id, it.Id == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Category, it.Category == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Name, it.Name == nil).
                Required(true).
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.PhotoUrls, it.PhotoUrls == nil).
                Required(true).
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Tags, it.Tags == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Status, it.Status == nil).
                
                
                EnumPattern(PetSTATUSPattern).
                Valid(factory)) {
        return false
    }

    return true
}

func (it *Pet) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil  {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *PetArray) Valid(factory swagger.ValidatorFactory) bool {
    for _, it := range *arr {
        if(!factory.NewValidator(it.Id, it.Id == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Category, it.Category == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Name, it.Name == nil).
        Required(true).
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.PhotoUrls, it.PhotoUrls == nil).
        Required(true).
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Tags, it.Tags == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Status, it.Status == nil).
        
        
        EnumPattern(PetSTATUSPattern).
        Valid(factory)) {
        return false
        }
    }
    return true
}

func (it *PetArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
    writer.WriteHeader(200)
    if err := producer.Produce(writer, it); err != nil  {
        panic(err) // let the recovery middleware deal with this
    }
}


func (it *Pet)this_is_call_dummy() {
    time.Now()
}




