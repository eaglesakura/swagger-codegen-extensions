package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"net/http"
	"strings"
)

type CreateUserParams struct {
	// Created user object
	Body *User
}

/*
Create user

This can only be done by the logged in user.
 param: Body Created user object
 return: void
*/
type CreateUserHandler func(context swagger.RequestContext, params *CreateUserParams) swagger.Responder

func (it *CreateUserParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Body, it.Body == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewCreateUserParams(binder swagger.RequestBinder) (*CreateUserParams, error) {
	result := &CreateUserParams{}

	if err := binder.BindBody("User", &result.Body); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type CreateUsersWithArrayInputParams struct {
	// List of user object
	Body *[]User
}

/*
Creates list of users with given input array


 param: Body List of user object
 return: void
*/
type CreateUsersWithArrayInputHandler func(context swagger.RequestContext, params *CreateUsersWithArrayInputParams) swagger.Responder

func (it *CreateUsersWithArrayInputParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Body, it.Body == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewCreateUsersWithArrayInputParams(binder swagger.RequestBinder) (*CreateUsersWithArrayInputParams, error) {
	result := &CreateUsersWithArrayInputParams{}

	if err := binder.BindBody("[]User", &result.Body); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type CreateUsersWithListInputParams struct {
	// List of user object
	Body *[]User
}

/*
Creates list of users with given input array


 param: Body List of user object
 return: void
*/
type CreateUsersWithListInputHandler func(context swagger.RequestContext, params *CreateUsersWithListInputParams) swagger.Responder

func (it *CreateUsersWithListInputParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Body, it.Body == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewCreateUsersWithListInputParams(binder swagger.RequestBinder) (*CreateUsersWithListInputParams, error) {
	result := &CreateUsersWithListInputParams{}

	if err := binder.BindBody("[]User", &result.Body); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type DeleteUserParams struct {
	// The name that needs to be deleted
	Username *string
}

/*
Delete user

This can only be done by the logged in user.
 param: Username The name that needs to be deleted
 return: void
*/
type DeleteUserHandler func(context swagger.RequestContext, params *DeleteUserParams) swagger.Responder

func (it *DeleteUserParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Username, it.Username == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewDeleteUserParams(binder swagger.RequestBinder) (*DeleteUserParams, error) {
	result := &DeleteUserParams{}

	if err := binder.BindPath("username", "string", &result.Username); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type GetUserByNameParams struct {
	// The name that needs to be fetched. Use user1 for testing.
	Username *string
}

/*
Get user by user name


 param: Username The name that needs to be fetched. Use user1 for testing.
 return: User
*/
type GetUserByNameHandler func(context swagger.RequestContext, params *GetUserByNameParams) swagger.Responder

func (it *GetUserByNameParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Username, it.Username == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewGetUserByNameParams(binder swagger.RequestBinder) (*GetUserByNameParams, error) {
	result := &GetUserByNameParams{}

	if err := binder.BindPath("username", "string", &result.Username); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type LoginUserParams struct {
	// The user name for login
	Username *string
	// The password for login in clear text
	Password *string
}

/*
Logs user into the system


 param: Username The user name for login
 param: Password The password for login in clear text
 return: string
*/
type LoginUserHandler func(context swagger.RequestContext, params *LoginUserParams) swagger.Responder

func (it *LoginUserParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Username, it.Username == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Password, it.Password == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewLoginUserParams(binder swagger.RequestBinder) (*LoginUserParams, error) {
	result := &LoginUserParams{}

	if err := binder.BindQuery("username", "string", &result.Username); err != nil {
		return nil, err
	}
	if err := binder.BindQuery("password", "string", &result.Password); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type LogoutUserParams struct {
}

/*
Logs out current logged in user session


 return: void
*/
type LogoutUserHandler func(context swagger.RequestContext, params *LogoutUserParams) swagger.Responder

func (it *LogoutUserParams) Valid(factory swagger.ValidatorFactory) bool {

	return true
}

// Bind from request
func NewLogoutUserParams(binder swagger.RequestBinder) (*LogoutUserParams, error) {
	result := &LogoutUserParams{}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type UpdateUserParams struct {
	// name that need to be updated
	Username *string
	// Updated user object
	Body *User
}

/*
Updated user

This can only be done by the logged in user.
 param: Username name that need to be updated
 param: Body Updated user object
 return: void
*/
type UpdateUserHandler func(context swagger.RequestContext, params *UpdateUserParams) swagger.Responder

func (it *UpdateUserParams) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Username, it.Username == nil).Required(true).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Body, it.Body == nil).Required(true).
		Valid(factory) {
		return false
	}

	return true
}

// Bind from request
func NewUpdateUserParams(binder swagger.RequestBinder) (*UpdateUserParams, error) {
	result := &UpdateUserParams{}

	if err := binder.BindPath("username", "string", &result.Username); err != nil {
		return nil, err
	}

	if err := binder.BindBody("User", &result.Body); err != nil {
		return nil, err
	}

	if !result.Valid(binder) {
		return nil, errors.New(400 /* Bad Request */, "Parameter validate error")
	}

	return result, nil
}

type UserApiController struct {
	CreateUser swagger.HandleRequest

	CreateUsersWithArrayInput swagger.HandleRequest

	CreateUsersWithListInput swagger.HandleRequest

	DeleteUser swagger.HandleRequest

	GetUserByName swagger.HandleRequest

	LoginUser swagger.HandleRequest

	LogoutUser swagger.HandleRequest

	UpdateUser swagger.HandleRequest
}

func NewUserApiController() *UserApiController {
	result := &UserApiController{}

	result.CreateUser.Path = "/v2/user"
	result.CreateUser.Method = strings.ToUpper("Post")
	result.HandleCreateUser(func(context swagger.RequestContext, params *CreateUserParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl CreateUser"))
	})

	result.CreateUsersWithArrayInput.Path = "/v2/user/createWithArray"
	result.CreateUsersWithArrayInput.Method = strings.ToUpper("Post")
	result.HandleCreateUsersWithArrayInput(func(context swagger.RequestContext, params *CreateUsersWithArrayInputParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl CreateUsersWithArrayInput"))
	})

	result.CreateUsersWithListInput.Path = "/v2/user/createWithList"
	result.CreateUsersWithListInput.Method = strings.ToUpper("Post")
	result.HandleCreateUsersWithListInput(func(context swagger.RequestContext, params *CreateUsersWithListInputParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl CreateUsersWithListInput"))
	})

	result.DeleteUser.Path = "/v2/user/{username}"
	result.DeleteUser.Method = strings.ToUpper("Delete")
	result.HandleDeleteUser(func(context swagger.RequestContext, params *DeleteUserParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl DeleteUser"))
	})

	result.GetUserByName.Path = "/v2/user/{username}"
	result.GetUserByName.Method = strings.ToUpper("Get")
	result.HandleGetUserByName(func(context swagger.RequestContext, params *GetUserByNameParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl GetUserByName"))
	})

	result.LoginUser.Path = "/v2/user/login"
	result.LoginUser.Method = strings.ToUpper("Get")
	result.HandleLoginUser(func(context swagger.RequestContext, params *LoginUserParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl LoginUser"))
	})

	result.LogoutUser.Path = "/v2/user/logout"
	result.LogoutUser.Method = strings.ToUpper("Get")
	result.HandleLogoutUser(func(context swagger.RequestContext, params *LogoutUserParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl LogoutUser"))
	})

	result.UpdateUser.Path = "/v2/user/{username}"
	result.UpdateUser.Method = strings.ToUpper("Put")
	result.HandleUpdateUser(func(context swagger.RequestContext, params *UpdateUserParams) swagger.Responder {
		return context.NewBindErrorResponse(errors.New(501, "Not Impl UpdateUser"))
	})

	return result
}

func (it *UserApiController) HandleCreateUser(handler CreateUserHandler) {
	it.CreateUser.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewCreateUserParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) HandleCreateUsersWithArrayInput(handler CreateUsersWithArrayInputHandler) {
	it.CreateUsersWithArrayInput.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewCreateUsersWithArrayInputParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) HandleCreateUsersWithListInput(handler CreateUsersWithListInputHandler) {
	it.CreateUsersWithListInput.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewCreateUsersWithListInputParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) HandleDeleteUser(handler DeleteUserHandler) {
	it.DeleteUser.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewDeleteUserParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) HandleGetUserByName(handler GetUserByNameHandler) {
	it.GetUserByName.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewGetUserByNameParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) HandleLoginUser(handler LoginUserHandler) {
	it.LoginUser.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewLoginUserParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) HandleLogoutUser(handler LogoutUserHandler) {
	it.LogoutUser.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewLogoutUserParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) HandleUpdateUser(handler UpdateUserHandler) {
	it.UpdateUser.HandlerFunc = func(context swagger.RequestContext, request *http.Request) swagger.Responder {
		binder, err := context.NewRequestBinder(request)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		params, err := NewUpdateUserParams(binder)
		if err != nil {
			return context.NewBindErrorResponse(err)
		}

		return handler(context, params)
	}
}

func (it *UserApiController) MapHandlers(mapper swagger.HandleMapper) {

	mapper.PutHandler(it.CreateUser)

	mapper.PutHandler(it.CreateUsersWithArrayInput)

	mapper.PutHandler(it.CreateUsersWithListInput)

	mapper.PutHandler(it.DeleteUser)

	mapper.PutHandler(it.GetUserByName)

	mapper.PutHandler(it.LoginUser)

	mapper.PutHandler(it.LogoutUser)

	mapper.PutHandler(it.UpdateUser)

}
