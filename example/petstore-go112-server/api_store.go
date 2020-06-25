package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"net/http"
	"strings"
)

type DeleteOrderParams struct {
	// ID of the order that needs to be deleted
	OrderId *int64
}

/*
Delete purchase order by ID

For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
 param: OrderId ID of the order that needs to be deleted
 return: void
*/
type DeleteOrderHandler func(context swagger.RequestContext, params *DeleteOrderParams) swagger.Responder

func (it *DeleteOrderParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.OrderId, it.OrderId == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewDeleteOrderParams(binder swagger.RequestBinder) (*DeleteOrderParams, error) {
	result := &DeleteOrderParams{}

	if err := binder.BindPath("orderId", "int64", &result.OrderId); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type GetInventoryParams struct {
}

/*
Returns pet inventories by status

Returns a map of status codes to quantities
 return: map[string]int32
*/
type GetInventoryHandler func(context swagger.RequestContext, params *GetInventoryParams) swagger.Responder

func (it *GetInventoryParams) Valid(factory swagger.ValidatorFactory) bool {

	return true
}

// Bind from request
func NewGetInventoryParams(binder swagger.RequestBinder) (*GetInventoryParams, error) {
	result := &GetInventoryParams{}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type GetOrderByIdParams struct {
	// ID of pet that needs to be fetched
	OrderId *int64
}

/*
Find purchase order by ID

For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions
 param: OrderId ID of pet that needs to be fetched
 return: Order
*/
type GetOrderByIdHandler func(context swagger.RequestContext, params *GetOrderByIdParams) swagger.Responder

func (it *GetOrderByIdParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.OrderId, it.OrderId == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewGetOrderByIdParams(binder swagger.RequestBinder) (*GetOrderByIdParams, error) {
	result := &GetOrderByIdParams{}

	if err := binder.BindPath("orderId", "int64", &result.OrderId); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type PlaceOrderParams struct {
	// order placed for purchasing the pet
	Body *Order
}

/*
Place an order for a pet


 param: Body order placed for purchasing the pet
 return: Order
*/
type PlaceOrderHandler func(context swagger.RequestContext, params *PlaceOrderParams) swagger.Responder

func (it *PlaceOrderParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Body, it.Body == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewPlaceOrderParams(binder swagger.RequestBinder) (*PlaceOrderParams, error) {
	result := &PlaceOrderParams{}

	if err := binder.BindBody("Order", &result.Body); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type StoreApiController struct {
	DeleteOrder swagger.HandleRequest

	GetInventory swagger.HandleRequest

	GetOrderById swagger.HandleRequest

	PlaceOrder swagger.HandleRequest
}

func NewStoreApiController() *StoreApiController {
	result := &StoreApiController{}

	result.DeleteOrder.Path = "/v2/store/order/{orderId}"
	result.DeleteOrder.Method = strings.ToUpper("Delete")
	result.HandleDeleteOrder(func(context swagger.RequestContext, params *DeleteOrderParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl DeleteOrder"))
	})

	result.GetInventory.Path = "/v2/store/inventory"
	result.GetInventory.Method = strings.ToUpper("Get")
	result.HandleGetInventory(func(context swagger.RequestContext, params *GetInventoryParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl GetInventory"))
	})

	result.GetOrderById.Path = "/v2/store/order/{orderId}"
	result.GetOrderById.Method = strings.ToUpper("Get")
	result.HandleGetOrderById(func(context swagger.RequestContext, params *GetOrderByIdParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl GetOrderById"))
	})

	result.PlaceOrder.Path = "/v2/store/order"
	result.PlaceOrder.Method = strings.ToUpper("Post")
	result.HandlePlaceOrder(func(context swagger.RequestContext, params *PlaceOrderParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl PlaceOrder"))
	})

	return result
}

func (it *StoreApiController) HandleDeleteOrder(handler DeleteOrderHandler) {
	it.DeleteOrder.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewDeleteOrderParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *StoreApiController) HandleGetInventory(handler GetInventoryHandler) {
	it.GetInventory.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewGetInventoryParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *StoreApiController) HandleGetOrderById(handler GetOrderByIdHandler) {
	it.GetOrderById.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewGetOrderByIdParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *StoreApiController) HandlePlaceOrder(handler PlaceOrderHandler) {
	it.PlaceOrder.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewPlaceOrderParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *StoreApiController) MapHandlers(mapper swagger.HandleMapper) {

	mapper.PutHandler(it.DeleteOrder)

	mapper.PutHandler(it.GetInventory)

	mapper.PutHandler(it.GetOrderById)

	mapper.PutHandler(it.PlaceOrder)

}
