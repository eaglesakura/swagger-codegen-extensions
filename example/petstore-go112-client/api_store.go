package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
    "fmt"
    "io"
    "net/url"
    "strings"
    "github.com/eaglesakura/swagger-go-core"
    "github.com/eaglesakura/swagger-go-core/errors"
    "github.com/eaglesakura/swagger-go-core/utils"
)


const StoreApi_BasePath string = "/v2"
type StoreApi struct {
    BasePath string
}


func NewStoreApi() *StoreApi {
    return &StoreApi{
        BasePath:StoreApi_BasePath,
    }
}

        /*
        Delete purchase order by ID
        For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
        */
        type StoreApiDeleteOrderRequest struct {
            /*
            ID of the order that needs to be deleted
            */
            orderId   *int64

        }

        /*
        Delete purchase order by ID
        For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors

          result: void
        */
        func (it *StoreApi)DeleteOrder(_client swagger.FetchClient, _request *StoreApiDeleteOrderRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.orderId, _request.orderId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'orderId' when calling DeleteOrder")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/store/order/{orderId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "orderId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.orderId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Delete"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Returns pet inventories by status
        Returns a map of status codes to quantities
        */
        type StoreApiGetInventoryRequest struct {
        }

        /*
        Returns pet inventories by status
        Returns a map of status codes to quantities

          result: map[string]int32
        */
        func (it *StoreApi)GetInventory(_client swagger.FetchClient, _request *StoreApiGetInventoryRequest, result *map[string]int32) error {

        // create path and map variables
        {
            localVarPath := strings.Replace("/store/inventory","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Find purchase order by ID
        For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions
        */
        type StoreApiGetOrderByIdRequest struct {
            /*
            ID of pet that needs to be fetched
            */
            orderId   *int64

        }

        /*
        Find purchase order by ID
        For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions

          result: Order
        */
        func (it *StoreApi)GetOrderById(_client swagger.FetchClient, _request *StoreApiGetOrderByIdRequest, result *Order) error {
            if(!_client.NewValidator(_request.orderId, _request.orderId == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'orderId' when calling GetOrderById")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/store/order/{orderId}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "orderId" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.orderId)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Place an order for a pet
        
        */
        type StoreApiPlaceOrderRequest struct {
            /*
            order placed for purchasing the pet
            */
            body   *Order

        }

        /*
        Place an order for a pet
        

          result: Order
        */
        func (it *StoreApi)PlaceOrder(_client swagger.FetchClient, _request *StoreApiPlaceOrderRequest, result *Order) error {
            if(!_client.NewValidator(_request.body, _request.body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'body' when calling PlaceOrder")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/store/order","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Post"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }

            if _request.body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.body))
            }

            return _client.Fetch(result)
        }


func (it *StoreApi)this_is_call_dummy() {
    v := url.Values{}
    v.Add("Key", "Value")

    errors.New(0, "stub")
    strings.ToUpper("")
    fmt.Sprintf("%v", "")
    io.ReadAtLeast(nil, nil, 0)
}
