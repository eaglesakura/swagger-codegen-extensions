package com.example.petstore.client

import com.example.petstore.client.* // ktlint-disable
import java.util.* // ktlint-disable
import java.io.* // ktlint-disable
import java.math.* // ktlint-disable
import retrofit2.Response // ktlint-disable
import okhttp3.ResponseBody // ktlint-disable
import okhttp3.OkHttpClient // ktlint-disable
import retrofit2.* // ktlint-disable
import retrofit2.http.* // ktlint-disable
import com.squareup.moshi.* // ktlint-disable
import retrofit2.converter.moshi.* // ktlint-disable

/**
 * UserApi Client.
 *
 * e.g.)
 * // build.gradle
 * dependencies {
 *      implementation "com.squareup.moshi:moshi:1.5.0"
 *      implementation "com.squareup.moshi:moshi-kotlin:1.5.0"
 *      implementation "com.squareup.retrofit2:retrofit:2.9.0"
 *      implementation "com.squareup.retrofit2:converter-moshi:2.4.0"
 * }
 *
 * // Example.kt
 * suspend fun callApiSample() {
 *      val exampleApi: UserApi = UserApi.create("https://example.com")
 *      exampleApi.fooApi()
 * }
 *
 * @link https://github.com/eaglesakura/swagger-codegen-extensions
 */
interface UserApi {
    /**
     * Create user
     * This can only be done by the logged in user.
     *
     * @param body Created user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @POST("/v2/user")
    suspend fun createUser(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: User
    ): Response<ResponseBody>

    /**
     * Create user
     * This can only be done by the logged in user.
     *
     * @param body Created user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @POST("/v2/user")
    suspend fun createUserBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: User
    ): Response<ResponseBody>

    /**
     * Create user
     * This can only be done by the logged in user.
     *
     * @param body Created user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @POST("/v2/user")
    suspend fun createUserStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: User
    ): Response<ResponseBody>
    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @POST("/v2/user/createWithArray")
    suspend fun createUsersWithArrayInput(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: kotlin.Array<User>
    ): Response<ResponseBody>

    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @POST("/v2/user/createWithArray")
    suspend fun createUsersWithArrayInputBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: kotlin.Array<User>
    ): Response<ResponseBody>

    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @POST("/v2/user/createWithArray")
    suspend fun createUsersWithArrayInputStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: kotlin.Array<User>
    ): Response<ResponseBody>
    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @POST("/v2/user/createWithList")
    suspend fun createUsersWithListInput(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: kotlin.Array<User>
    ): Response<ResponseBody>

    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @POST("/v2/user/createWithList")
    suspend fun createUsersWithListInputBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: kotlin.Array<User>
    ): Response<ResponseBody>

    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @POST("/v2/user/createWithList")
    suspend fun createUsersWithListInputStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: kotlin.Array<User>
    ): Response<ResponseBody>
    /**
     * Delete user
     * This can only be done by the logged in user.
     *
     * @param username The name that needs to be deleted/ example :: username_example
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @DELETE("/v2/user/{username}")
    suspend fun deleteUser(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String
    ): Response<ResponseBody>

    /**
     * Delete user
     * This can only be done by the logged in user.
     *
     * @param username The name that needs to be deleted/ example :: username_example
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @DELETE("/v2/user/{username}")
    suspend fun deleteUserBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String
    ): Response<ResponseBody>

    /**
     * Delete user
     * This can only be done by the logged in user.
     *
     * @param username The name that needs to be deleted/ example :: username_example
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @DELETE("/v2/user/{username}")
    suspend fun deleteUserStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String
    ): Response<ResponseBody>
    /**
     * Get user by user name
     * 
     *
     * @param username The name that needs to be fetched. Use user1 for testing. / example :: username_example
     * @return User
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @GET("/v2/user/{username}")
    suspend fun getUserByName(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String
    ): Response<User>

    /**
     * Get user by user name
     * 
     *
     * @param username The name that needs to be fetched. Use user1 for testing. / example :: username_example
     * @return User
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @GET("/v2/user/{username}")
    suspend fun getUserByNameBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String
    ): Response<ResponseBody>

    /**
     * Get user by user name
     * 
     *
     * @param username The name that needs to be fetched. Use user1 for testing. / example :: username_example
     * @return User
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @GET("/v2/user/{username}")
    suspend fun getUserByNameStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String
    ): Response<ResponseBody>
    /**
     * Logs user into the system
     * 
     *
     * @param username The user name for login/ example :: username_example
     * @param password The password for login in clear text/ example :: password_example
     * @return kotlin.String
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @GET("/v2/user/login")
    suspend fun loginUser(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Query("username")
            username: kotlin.String, 
            
            @Query("password")
            password: kotlin.String
    ): Response<kotlin.String>

    /**
     * Logs user into the system
     * 
     *
     * @param username The user name for login/ example :: username_example
     * @param password The password for login in clear text/ example :: password_example
     * @return kotlin.String
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @GET("/v2/user/login")
    suspend fun loginUserBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Query("username")
            username: kotlin.String, 
            
            @Query("password")
            password: kotlin.String
    ): Response<ResponseBody>

    /**
     * Logs user into the system
     * 
     *
     * @param username The user name for login/ example :: username_example
     * @param password The password for login in clear text/ example :: password_example
     * @return kotlin.String
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @GET("/v2/user/login")
    suspend fun loginUserStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Query("username")
            username: kotlin.String, 
                        @Query("password")
            password: kotlin.String
    ): Response<ResponseBody>
    /**
     * Logs out current logged in user session
     * 
     *
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @GET("/v2/user/logout")
    suspend fun logoutUser(
            @HeaderMap customHeaders: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Logs out current logged in user session
     * 
     *
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @GET("/v2/user/logout")
    suspend fun logoutUserBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Logs out current logged in user session
     * 
     *
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @GET("/v2/user/logout")
    suspend fun logoutUserStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Updated user
     * This can only be done by the logged in user.
     *
     * @param username name that need to be updated/ example :: username_example
     * @param body Updated user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @PUT("/v2/user/{username}")
    suspend fun updateUser(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String, 
            
            @Body 
            body: User
    ): Response<ResponseBody>

    /**
     * Updated user
     * This can only be done by the logged in user.
     *
     * @param username name that need to be updated/ example :: username_example
     * @param body Updated user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @PUT("/v2/user/{username}")
    suspend fun updateUserBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String, 
            
            @Body 
            body: User
    ): Response<ResponseBody>

    /**
     * Updated user
     * This can only be done by the logged in user.
     *
     * @param username name that need to be updated/ example :: username_example
     * @param body Updated user object
     * @return void
     * @link https://github.com/eaglesakura/swagger-codegen-extensions
     */
    @Streaming
    @PUT("/v2/user/{username}")
    suspend fun updateUserStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Header("username")
            username: kotlin.String, 
                        @Body 
            body: User
    ): Response<ResponseBody>
    companion object {
        /**
         * Default UserApi Factory.
         *
         * @link https://github.com/eaglesakura/swagger-codegen-extensions
         */
        fun create(baseUrl: String, okHttpClient: OkHttpClient? = null, moshi: Moshi = InternalUtils.moshi, block: (builder: Retrofit.Builder)->Retrofit.Builder = { it }): UserApi {
            val url = if(baseUrl.endsWith("/")) {
                baseUrl
            } else {
                "$baseUrl/"
            }
            val builder = Retrofit.Builder()
                                .baseUrl(url)
                                .addConverterFactory(MoshiConverterFactory.create(moshi))
            if(okHttpClient != null) {
                builder.client(okHttpClient)
            }
            return  block(builder).build().create(UserApi::class.java)
        }
    }
}

