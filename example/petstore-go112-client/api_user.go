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
            Body   *User

        }

        /*
        Create user
        This can only be done by the logged in user.

          result: void
        */
        func (it *UserApi)CreateUser(_client swagger.FetchClient, _request *UserApiCreateUserRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.Body, _request.Body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Body' when calling CreateUser")
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

            if _request.Body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.Body))
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
            Body   *[]User

        }

        /*
        Creates list of users with given input array
        

          result: void
        */
        func (it *UserApi)CreateUsersWithArrayInput(_client swagger.FetchClient, _request *UserApiCreateUsersWithArrayInputRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.Body, _request.Body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Body' when calling CreateUsersWithArrayInput")
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

            if _request.Body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.Body))
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
            Body   *[]User

        }

        /*
        Creates list of users with given input array
        

          result: void
        */
        func (it *UserApi)CreateUsersWithListInput(_client swagger.FetchClient, _request *UserApiCreateUsersWithListInputRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.Body, _request.Body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Body' when calling CreateUsersWithListInput")
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

            if _request.Body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.Body))
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
            Username   *string

        }

        /*
        Delete user
        This can only be done by the logged in user.

          result: void
        */
        func (it *UserApi)DeleteUser(_client swagger.FetchClient, _request *UserApiDeleteUserRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.Username, _request.Username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Username' when calling DeleteUser")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/{username}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "username" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.Username)), -1)
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
            Username   *string

        }

        /*
        Get user by user name
        

          result: User
        */
        func (it *UserApi)GetUserByName(_client swagger.FetchClient, _request *UserApiGetUserByNameRequest, result *User) error {
            if(!_client.NewValidator(_request.Username, _request.Username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Username' when calling GetUserByName")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/{username}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "username" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.Username)), -1)
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
            Username   *string

            /*
            The password for login in clear text
            */
            Password   *string

        }

        /*
        Logs user into the system
        

          result: string
        */
        func (it *UserApi)LoginUser(_client swagger.FetchClient, _request *UserApiLoginUserRequest, result *string) error {
            if(!_client.NewValidator(_request.Username, _request.Username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Username' when calling LoginUser")
            }
            if(!_client.NewValidator(_request.Password, _request.Password == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Password' when calling LoginUser")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/login","{format}","json", -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Get"))
        }

            if _request.Username != nil {
                _client.AddQueryParam("username", utils.ParameterToString(_request.Username))
            }
            if _request.Password != nil {
                _client.AddQueryParam("password", utils.ParameterToString(_request.Password))
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
            Username   *string

            /*
            Updated user object
            */
            Body   *User

        }

        /*
        Updated user
        This can only be done by the logged in user.

          result: void
        */
        func (it *UserApi)UpdateUser(_client swagger.FetchClient, _request *UserApiUpdateUserRequest, result interface{} ) error {
            if(!_client.NewValidator(_request.Username, _request.Username == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Username' when calling UpdateUser")
            }
            if(!_client.NewValidator(_request.Body, _request.Body == nil).Required(true).Valid(_client)) {
                return errors.New(0, "Missing the required parameter 'Body' when calling UpdateUser")
            }

        // create path and map variables
        {
            localVarPath := strings.Replace("/user/{username}","{format}","json", -1)
                localVarPath = strings.Replace(localVarPath, "{" + "username" + "}", utils.EscapeString(fmt.Sprintf("%v", *_request.Username)), -1)
            _client.SetApiPath(utils.AddPath(it.BasePath, localVarPath))
            _client.SetMethod(strings.ToUpper("Put"))
        }



        localVarFormParams := url.Values{}
        formEnable := false
        

        if formEnable {
            _client.SetPayload(utils.NewFormPayload(localVarFormParams))
        }

            if _request.Body != nil {
                _client.SetPayload(utils.NewJsonPayload(_request.Body))
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
