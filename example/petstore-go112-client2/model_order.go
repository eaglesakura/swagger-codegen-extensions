package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"time"
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

func (it OrderSTATUS) Ptr() *OrderSTATUS {
	return &it
}
func (it *OrderSTATUS) Value() OrderSTATUS {
	return *it
}
func (it *OrderSTATUS) Valid() bool {
	if it == nil {
		return false
	}
	value := string(*it)
	for _, v := range OrderSTATUSPattern {
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

// Update self from builder callback.
func (it *Order) Update(builder func(ref *Order)) *Order {
	builder(it)
	return it
}

// Copy and Update self from builder callback.
func (it *Order) Copy(builder func(copied *Order)) *Order {
	result := *it
	builder(&result)
	return &result
}

// Set Id
func (it *Order) SetId(newId int64) {
	it.Id = &newId
}

// Require value of Id
func (it *Order) RequireId() int64 {
	if it.Id == nil {
		panic(xerrors.Errorf("Order.Id is nil"))
	}
	return *it.Id
}

// Get value of Id / or default
func (it *Order) GetId() int64 {
	if it.Id != nil {
		return *it.Id
	}
	result := new(int64)
	return *result
}

// Set PetId
func (it *Order) SetPetId(newPetId int64) {
	it.PetId = &newPetId
}

// Require value of PetId
func (it *Order) RequirePetId() int64 {
	if it.PetId == nil {
		panic(xerrors.Errorf("Order.PetId is nil"))
	}
	return *it.PetId
}

// Get value of PetId / or default
func (it *Order) GetPetId() int64 {
	if it.PetId != nil {
		return *it.PetId
	}
	result := new(int64)
	return *result
}

// Set Quantity
func (it *Order) SetQuantity(newQuantity int32) {
	it.Quantity = &newQuantity
}

// Require value of Quantity
func (it *Order) RequireQuantity() int32 {
	if it.Quantity == nil {
		panic(xerrors.Errorf("Order.Quantity is nil"))
	}
	return *it.Quantity
}

// Get value of Quantity / or default
func (it *Order) GetQuantity() int32 {
	if it.Quantity != nil {
		return *it.Quantity
	}
	result := new(int32)
	return *result
}

// Set ShipDate
func (it *Order) SetShipDate(newShipDate time.Time) {
	it.ShipDate = &newShipDate
}

// Require value of ShipDate
func (it *Order) RequireShipDate() time.Time {
	if it.ShipDate == nil {
		panic(xerrors.Errorf("Order.ShipDate is nil"))
	}
	return *it.ShipDate
}

// Get value of ShipDate / or default
func (it *Order) GetShipDate() time.Time {
	if it.ShipDate != nil {
		return *it.ShipDate
	}
	result := new(time.Time)
	return *result
}

// Set Status
func (it *Order) SetStatus(newStatus OrderSTATUS) {
	it.Status = &newStatus
}

// Require value of Status
func (it *Order) RequireStatus() OrderSTATUS {
	if it.Status == nil {
		panic(xerrors.Errorf("Order.Status is nil"))
	}
	return *it.Status
}

// Get value of Status / or default
func (it *Order) GetStatus() OrderSTATUS {
	if it.Status != nil {
		return *it.Status
	}
	result := new(OrderSTATUS)
	return *result
}

// Set Complete
func (it *Order) SetComplete(newComplete bool) {
	it.Complete = &newComplete
}

// Require value of Complete
func (it *Order) RequireComplete() bool {
	if it.Complete == nil {
		panic(xerrors.Errorf("Order.Complete is nil"))
	}
	return *it.Complete
}

// Get value of Complete / or default
func (it *Order) GetComplete() bool {
	if it.Complete != nil {
		return *it.Complete
	}
	result := new(bool)
	return *result
}

// encode to json
func (it Order) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}
func (it Order) Json() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

type OrderArray []Order

func (it *Order) Valid() error {
	if it.Status != nil && !it.Status.Valid() {
		return xerrors.Errorf("'Order.Status' is invalid")
	}

	return nil
}

func (it *Order) this_is_call_dummy() {
	time.Now()
	xerrors.Errorf("")
}
