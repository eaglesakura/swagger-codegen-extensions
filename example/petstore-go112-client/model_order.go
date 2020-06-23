package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions



import (
    "time"
    "net/http"
    "encoding/json"
    "github.com/eaglesakura/swagger-go-core"
)


type OrderSTATUS string
type OrderSTATUSArray []OrderSTATUS

const OrderSTATUS_placed OrderSTATUS = OrderSTATUS("placed")
const OrderSTATUS_approved OrderSTATUS = OrderSTATUS("approved")
const OrderSTATUS_delivered OrderSTATUS = OrderSTATUS("delivered")
var OrderSTATUSPattern []string = []string{

    "placed", 
    "approved", 
    "delivered", 
}
func (it OrderSTATUS)Ptr() *OrderSTATUS {
    return &it
}
func (it *OrderSTATUS)Value() OrderSTATUS {
    return *it
}
func (it *OrderSTATUS)Valid(pattern []string) bool {
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
type Order struct {

	Id *int64 `json:"id,omitempty"`

	PetId *int64 `json:"petId,omitempty"`

	Quantity *int32 `json:"quantity,omitempty"`

	ShipDate *time.Time `json:"shipDate,omitempty"`

	// Order Status
	Status *OrderSTATUS `json:"status,omitempty"`

	Complete *bool `json:"complete,omitempty"`
}

// encode to json
func (it Order)String() string {
    buf, _ := json.Marshal(it)
    return string(buf)
}
func (it Order)Json() string {
    buf, _ := json.Marshal(it)
    return string(buf)
}

type OrderArray []Order

func (it *Order) Valid(factory swagger.ValidatorFactory) bool {
    if(!factory.NewValidator(it.Id, it.Id == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.PetId, it.PetId == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Quantity, it.Quantity == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.ShipDate, it.ShipDate == nil).
                
                
                
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Status, it.Status == nil).
                
                
                EnumPattern(OrderSTATUSPattern).
                Valid(factory)) {
        return false
    }
    if(!factory.NewValidator(it.Complete, it.Complete == nil).
                
                
                
                Valid(factory)) {
        return false
    }

    return true
}

func (it *Order) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil  {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *OrderArray) Valid(factory swagger.ValidatorFactory) bool {
    for _, it := range *arr {
        if(!factory.NewValidator(it.Id, it.Id == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.PetId, it.PetId == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Quantity, it.Quantity == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.ShipDate, it.ShipDate == nil).
        
        
        
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Status, it.Status == nil).
        
        
        EnumPattern(OrderSTATUSPattern).
        Valid(factory)) {
        return false
        }
        if(!factory.NewValidator(it.Complete, it.Complete == nil).
        
        
        
        Valid(factory)) {
        return false
        }
    }
    return true
}

func (it *OrderArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
    writer.WriteHeader(200)
    if err := producer.Produce(writer, it); err != nil  {
        panic(err) // let the recovery middleware deal with this
    }
}


func (it *Order)this_is_call_dummy() {
    time.Now()
}




