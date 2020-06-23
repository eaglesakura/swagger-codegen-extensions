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


const UserApi_BasePath string = "/v2"
type UserApi struct {
    BasePath string
}


func NewUserApi() *UserApi {
    return &UserApi{
        BasePath:UserApi_BasePath,
    }
}

        /*
        Create user
        This can only be done by the logged in user.
        */
        type UserApiCreateUserRequest struct {
            /*
            Created user object
            */
            body   *User

        }

        /*
        Create user
        This can only be done by the logged in user.

          result: void
        */
        func (it *UserApi)CreateUser(_client swagger.FetchClient, _request *UserApiCreateUserRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.body, _request.body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'body' when calling CreateUser")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user","{format}","json", -1)
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
        /*
        Creates list of users with given input array
        
        */
        type UserApiCreateUsersWithArrayInputRequest struct {
            /*
            List of user object
            */
            body   *[]User

        }

        /*
        Creates list of users with given input array
        

          result: void
        */
        func (it *UserApi)CreateUsersWithArrayInput(_client swagger.FetchClient, _request *UserApiCreateUsersWithArrayInputRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.body, _request.body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'body' when calling CreateUsersWithArrayInput")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/createWithArray","{format}","json", -1)
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
        /*
        Creates list of users with given input array
        
        */
        type UserApiCreateUsersWithListInputRequest struct {
            /*
            List of user object
            */
            body   *[]User

        }

        /*
        Creates list of users with given input array
        

          result: void
        */
        func (it *UserApi)CreateUsersWithListInput(_client swagger.FetchClient, _request *UserApiCreateUsersWithListInputRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.body, _request.body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'body' when calling CreateUsersWithListInput")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/createWithList","{format}","json", -1)
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
        /*
        Delete user
        This can only be done by the logged in user.
        */
        type UserApiDeleteUserRequest struct {
            /*
            The name that needs to be deleted
            */
            username   *string

        }

        /*
        Delete user
        This can only be done by the logged in user.

          result: void
        */
        func (it *UserApi)DeleteUser(_client swagger.FetchClient, _request *UserApiDeleteUserRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.username, _request.username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'username' when calling DeleteUser")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/{username}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "username" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.username)), -1)
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
        Get user by user name
        
        */
        type UserApiGetUserByNameRequest struct {
            /*
            The name that needs to be fetched. Use user1 for testing. 
            */
            username   *string

        }

        /*
        Get user by user name
        

          result: User
        */
        func (it *UserApi)GetUserByName(_client swagger.FetchClient, _request *UserApiGetUserByNameRequest, result *User) error {
            if(!_client.NewValidator(_request.username, _request.username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'username' when calling GetUserByName")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/{username}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "username" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.username)), -1)
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
        Logs user into the system
        
        */
        type UserApiLoginUserRequest struct {
            /*
            The user name for login
            */
            username   *string

            /*
            The password for login in clear text
            */
            password   *string

        }

        /*
        Logs user into the system
        

          result: string
        */
        func (it *UserApi)LoginUser(_client swagger.FetchClient, _request *UserApiLoginUserRequest, result *string) error {
            if(!_client.NewValidator(_request.username, _request.username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'username' when calling LoginUser")
            }
            if(!_client.NewValidator(_request.password, _request.password == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'password' when calling LoginUser")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/login","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }

            if _request.username != nil {
                _client.AddQueryParam("username", utils.ParameterToString(_request.username))
            }
            if _request.password != nil {
                _client.AddQueryParam("password", utils.ParameterToString(_request.password))
            }


        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }


            return _client.Fetch(result)
        }
        /*
        Logs out current logged in user session
        
        */
        type UserApiLogoutUserRequest struct {
        }

        /*
        Logs out current logged in user session
        

          result: void
        */
        func (it *UserApi)LogoutUser(_client swagger.FetchClient, _request *UserApiLogoutUserRequest, result interface{} ) error {

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/logout","{format}","json", -1)
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
        Updated user
        This can only be done by the logged in user.
        */
        type UserApiUpdateUserRequest struct {
            /*
            name that need to be updated
            */
            username   *string

            /*
            Updated user object
            */
            body   *User

        }

        /*
        Updated user
        This can only be done by the logged in user.

          result: void
        */
        func (it *UserApi)UpdateUser(_client swagger.FetchClient, _request *UserApiUpdateUserRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.username, _request.username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'username' when calling UpdateUser")
            }
            if(!_client.NewValidator(_request.body, _request.body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'body' when calling UpdateUser")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/{username}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "username" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.username)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Put"))
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


func (it *UserApi)this_is_call_dummy() {
    v := url.Values{}
    v.Add("Key", "Value")

    errors.New(0, "stub")
    strings.ToUpper("")
    fmt.Sprintf("%v", "")
    io.ReadAtLeast(nil, nil, 0)
}
