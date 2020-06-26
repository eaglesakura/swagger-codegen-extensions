package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"encoding/json"
	"golang.org/x/xerrors"
	"net/http"
	"time"
)

type OrderStatus string
type OrderStatusArray []OrderStatus

const OrderStatus_placed = OrderStatus("placed")
const OrderStatus_approved = OrderStatus("approved")
const OrderStatus_delivered = OrderStatus("delivered")

func (it OrderStatus) String() string {
	return string(it)
}

// Pattern for OrderStatus
var OrderStatusValues = []OrderStatus{

	OrderStatus_placed,

	OrderStatus_approved,

	OrderStatus_delivered,
}

func (it *OrderStatus) Valid() error {
	if it == nil {
		return xerrors.Errorf("OrderStatus(nil)")
	}
	switch *it {

	case OrderStatus_placed:
		return nil

	case OrderStatus_approved:
		return nil

	case OrderStatus_delivered:
		return nil

	}
	return xerrors.Errorf("invalid OrderStatus(%v)", string(*it))
}

//
// Order Interface for alternative definitions.
type IOrder interface {
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

	// Remove PetId property
	RemovePetId()

	// Get PetId property or default value.
	GetPetId() int64

	// Get PetId property or panic().
	RequirePetId() int64

	// Set PetId property.
	SetPetId(newPetId int64)

	// Remove Quantity property
	RemoveQuantity()

	// Get Quantity property or default value.
	GetQuantity() int32

	// Get Quantity property or panic().
	RequireQuantity() int32

	// Set Quantity property.
	SetQuantity(newQuantity int32)

	// Remove ShipDate property
	RemoveShipDate()

	// Get ShipDate property or default value.
	GetShipDate() time.Time

	// Get ShipDate property or panic().
	RequireShipDate() time.Time

	// Set ShipDate property.
	SetShipDate(newShipDate time.Time)

	// Remove Status property
	RemoveStatus()

	// Get Status property or default value.
	GetStatus() OrderStatus

	// Get Status property or panic().
	RequireStatus() OrderStatus

	// Set Status property.
	SetStatus(newStatus OrderStatus)

	// Remove Complete property
	RemoveComplete()

	// Get Complete property or default value.
	GetComplete() bool

	// Get Complete property or panic().
	RequireComplete() bool

	// Set Complete property.
	SetComplete(newComplete bool)
}

//
type Order struct {
	//
	Id *int64 `json:"id,omitempty"`
	//
	PetId *int64 `json:"petId,omitempty"`
	//
	Quantity *int32 `json:"quantity,omitempty"`
	//
	ShipDate *time.Time `json:"shipDate,omitempty"`
	// Order Status
	Status *OrderStatus `json:"status,omitempty"`
	//
	Complete *bool `json:"complete,omitempty"`
}
type OrderArray []Order

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

// DeepCopy to result ptr.
func (it *Order) MarshalCopy(result interface{}) error {
	body, _ := json.Marshal(it)
	err := json.Unmarshal(body, result)
	if err != nil {
		return xerrors.Errorf("Order.MarshalCopy failed, to='%v': %w", result, err)
	}
	return nil
}

// Remove Id
func (it *Order) RemoveId() {
	it.Id = nil
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

// Remove PetId
func (it *Order) RemovePetId() {
	it.PetId = nil
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

// Remove Quantity
func (it *Order) RemoveQuantity() {
	it.Quantity = nil
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

// Remove ShipDate
func (it *Order) RemoveShipDate() {
	it.ShipDate = nil
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

// Remove Status
func (it *Order) RemoveStatus() {
	it.Status = nil
}

// Set Status
func (it *Order) SetStatus(newStatus OrderStatus) {
	it.Status = &newStatus
}

// Require value of Status
func (it *Order) RequireStatus() OrderStatus {
	if it.Status == nil {
		panic(xerrors.Errorf("Order.Status is nil"))
	}
	return *it.Status
}

// Get value of Status / or default
func (it *Order) GetStatus() OrderStatus {
	if it.Status != nil {
		return *it.Status
	}
	result := new(OrderStatus)
	return *result
}

// Remove Complete
func (it *Order) RemoveComplete() {
	it.Complete = nil
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
func (it *Order) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *Order) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it *Order) Valid() error {
	if it.Id != nil {
		if err := validationValue(it.Id); err != nil {
			return xerrors.Errorf("'Order.Id' validation error: %w", err)
		}
	}
	if it.PetId != nil {
		if err := validationValue(it.PetId); err != nil {
			return xerrors.Errorf("'Order.PetId' validation error: %w", err)
		}
	}
	if it.Quantity != nil {
		if err := validationValue(it.Quantity); err != nil {
			return xerrors.Errorf("'Order.Quantity' validation error: %w", err)
		}
	}
	if it.ShipDate != nil {
		if err := validationValue(it.ShipDate); err != nil {
			return xerrors.Errorf("'Order.ShipDate' validation error: %w", err)
		}
	}
	if it.Status != nil {
		if err := validationValue(it.Status); err != nil {
			return xerrors.Errorf("'Order.Status' validation error: %w", err)
		}
	}
	if it.Complete != nil {
		if err := validationValue(it.Complete); err != nil {
			return xerrors.Errorf("'Order.Complete' validation error: %w", err)
		}
	}

	return nil
}

func (it OrderArray) Write(writer http.ResponseWriter, request *http.Request) {
	buf, _ := json.Marshal(it)
	_, _ = writer.Write(buf)
}

func (it OrderArray) Json() []byte {
	buf, _ := json.Marshal(it)
	return buf
}

func (it *Order) compilerDummy() {
	time.Now()
	_ = xerrors.Errorf("")

	var model Order
	var iModel IOrder = &model
	var swaggerModelRef swaggerModel = &model
	var swaggerResponseRef SwaggerResponse = &model
	var swaggerValidatableRef swaggerValidatable = &model

	var modelArray OrderArray
	swaggerResponseRef = modelArray

	iModel = iModel
	swaggerModelRef = swaggerModelRef
	swaggerResponseRef = swaggerResponseRef
	swaggerValidatableRef = swaggerValidatableRef
}
