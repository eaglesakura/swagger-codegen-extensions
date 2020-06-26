package petstore

// link: https://github.com/eaglesakura/swagger-codegen-extensions

import (
    "net/http"
)

// petstore
type UserApi struct {
    // API Endpoint
    // e.g.) "https://example.com/"
    Endpoint string

    // http.Client prefetch intercept function.
    Intercept func(client *http.Client, request *http.Request) (*http.Client, *http.Request)
}

// Update self from builder callback.
func (it *UserApi)Update(builder func(ref *UserApi)) *UserApi {
    builder(it)
    return it
}


// Copy and Update self from builder callback.
func (it *UserApi)Copy(builder func(copied *UserApi)) *UserApi {
    result := *it
    builder(&result)
    return &result
}

func NewUserApi() *UserApi {
    return &UserApi{
        Endpoint: "",
    }
}

func (it *UserApi)CreateUserPost(builder func(*UserApiCreateUserPostRequest)) *UserApiCreateUserPostRequest {
    result := &UserApiCreateUserPostRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
func (it *UserApi)CreateUsersWithArrayInputPost(builder func(*UserApiCreateUsersWithArrayInputPostRequest)) *UserApiCreateUsersWithArrayInputPostRequest {
    result := &UserApiCreateUsersWithArrayInputPostRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
func (it *UserApi)CreateUsersWithListInputPost(builder func(*UserApiCreateUsersWithListInputPostRequest)) *UserApiCreateUsersWithListInputPostRequest {
    result := &UserApiCreateUsersWithListInputPostRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
func (it *UserApi)DeleteUserDelete(builder func(*UserApiDeleteUserDeleteRequest)) *UserApiDeleteUserDeleteRequest {
    result := &UserApiDeleteUserDeleteRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
func (it *UserApi)GetUserByNameGet(builder func(*UserApiGetUserByNameGetRequest)) *UserApiGetUserByNameGetRequest {
    result := &UserApiGetUserByNameGetRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
func (it *UserApi)LoginUserGet(builder func(*UserApiLoginUserGetRequest)) *UserApiLoginUserGetRequest {
    result := &UserApiLoginUserGetRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
func (it *UserApi)LogoutUserGet(builder func(*UserApiLogoutUserGetRequest)) *UserApiLogoutUserGetRequest {
    result := &UserApiLogoutUserGetRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
func (it *UserApi)UpdateUserPut(builder func(*UserApiUpdateUserPutRequest)) *UserApiUpdateUserPutRequest {
    result := &UserApiUpdateUserPutRequest {
        Endpoint:   it.Endpoint,
        Intercept:  it.Intercept,
    }
    builder(result)
    return result
}
