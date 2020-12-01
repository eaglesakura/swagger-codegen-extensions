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
     */
    @POST("/v2/user")
    suspend fun createUser(
        
        
        
        @Body body: User,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Create user
     * This can only be done by the logged in user.
     *
     * @param body Created user object
     * @return void
     */
    @Streaming
    @POST("/v2/user")
    suspend fun createUserStreaming(
        
        
        
        @Body body: User,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     */
    @POST("/v2/user/createWithArray")
    suspend fun createUsersWithArrayInput(
        
        
        
        @Body body: kotlin.Array<User>,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     */
    @Streaming
    @POST("/v2/user/createWithArray")
    suspend fun createUsersWithArrayInputStreaming(
        
        
        
        @Body body: kotlin.Array<User>,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     */
    @POST("/v2/user/createWithList")
    suspend fun createUsersWithListInput(
        
        
        
        @Body body: kotlin.Array<User>,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Creates list of users with given input array
     * 
     *
     * @param body List of user object
     * @return void
     */
    @Streaming
    @POST("/v2/user/createWithList")
    suspend fun createUsersWithListInputStreaming(
        
        
        
        @Body body: kotlin.Array<User>,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Delete user
     * This can only be done by the logged in user.
     *
     * @param username The name that needs to be deleted/ example :: username_example
     * @return void
     */
    @DELETE("/v2/user/{username}")
    suspend fun deleteUser(
        
        @Path("username") username: kotlin.String,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Delete user
     * This can only be done by the logged in user.
     *
     * @param username The name that needs to be deleted/ example :: username_example
     * @return void
     */
    @Streaming
    @DELETE("/v2/user/{username}")
    suspend fun deleteUserStreaming(
        
        @Path("username") username: kotlin.String,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Get user by user name
     * 
     *
     * @param username The name that needs to be fetched. Use user1 for testing. / example :: username_example
     * @return User
     */
    @GET("/v2/user/{username}")
    suspend fun getUserByName(
        
        @Path("username") username: kotlin.String,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<User>

    /**
     * Get user by user name
     * 
     *
     * @param username The name that needs to be fetched. Use user1 for testing. / example :: username_example
     * @return User
     */
    @Streaming
    @GET("/v2/user/{username}")
    suspend fun getUserByNameStreaming(
        
        @Path("username") username: kotlin.String,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Logs user into the system
     * 
     *
     * @param username The user name for login/ example :: username_example
     * @param password The password for login in clear text/ example :: password_example
     * @return kotlin.String
     */
    @GET("/v2/user/login")
    suspend fun loginUser(
        
        
        @Query("username") username: kotlin.String,@Query("password") password: kotlin.String,
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<kotlin.String>

    /**
     * Logs user into the system
     * 
     *
     * @param username The user name for login/ example :: username_example
     * @param password The password for login in clear text/ example :: password_example
     * @return kotlin.String
     */
    @Streaming
    @GET("/v2/user/login")
    suspend fun loginUserStreaming(
        
        
        @Query("username") username: kotlin.String,@Query("password") password: kotlin.String,
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Logs out current logged in user session
     * 
     *
     * @return void
     */
    @GET("/v2/user/logout")
    suspend fun logoutUser(
        
        
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Logs out current logged in user session
     * 
     *
     * @return void
     */
    @Streaming
    @GET("/v2/user/logout")
    suspend fun logoutUserStreaming(
        
        
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Updated user
     * This can only be done by the logged in user.
     *
     * @param username name that need to be updated/ example :: username_example
     * @param body Updated user object
     * @return void
     */
    @PUT("/v2/user/{username}")
    suspend fun updateUser(
        
        @Path("username") username: kotlin.String,
        
        @Body body: User,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Updated user
     * This can only be done by the logged in user.
     *
     * @param username name that need to be updated/ example :: username_example
     * @param body Updated user object
     * @return void
     */
    @Streaming
    @PUT("/v2/user/{username}")
    suspend fun updateUserStreaming(
        
        @Path("username") username: kotlin.String,
        
        @Body body: User,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    companion object {
        /**
         * Default UserApi Factory.
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

