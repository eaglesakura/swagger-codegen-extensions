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
 * PetApi Client.
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
 *      val exampleApi: PetApi = PetApi.create("https://example.com")
 *      exampleApi.fooApi()
 * }
 *
 * @link https://github.com/eaglesakura/swagger-codegen-extensions
 */
interface PetApi {
    /**
     * Add a new pet to the store
     * 
     *
     * @param body Pet object that needs to be added to the store
     * @return void
     */
    @POST("/v2/pet")
    suspend fun addPet(
        
        
        
        @Body body: Pet,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Add a new pet to the store
     * 
     *
     * @param body Pet object that needs to be added to the store
     * @return void
     */
    @Streaming
    @POST("/v2/pet")
    suspend fun addPetStreaming(
        
        
        
        @Body body: Pet,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Deletes a pet
     * 
     *
     * @param petId Pet id to delete/ example :: 789
     * @param apiKey / example :: apiKey_example
     * @return void
     */
    @DELETE("/v2/pet/{petId}")
    suspend fun deletePet(
        @Header("api_key") apiKey: kotlin.String? = null,
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Deletes a pet
     * 
     *
     * @param petId Pet id to delete/ example :: 789
     * @param apiKey / example :: apiKey_example
     * @return void
     */
    @Streaming
    @DELETE("/v2/pet/{petId}")
    suspend fun deletePetStreaming(
        @Header("api_key") apiKey: kotlin.String? = null,
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Finds Pets by status
     * Multiple status values can be provided with comma separated strings
     *
     * @param status Status values that need to be considered for filter
     * @return kotlin.Array<Pet>
     */
    @GET("/v2/pet/findByStatus")
    suspend fun findPetsByStatus(
        
        
        @Query("status") status: kotlin.Array<kotlin.String>,
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<kotlin.Array<Pet>>

    /**
     * Finds Pets by status
     * Multiple status values can be provided with comma separated strings
     *
     * @param status Status values that need to be considered for filter
     * @return kotlin.Array<Pet>
     */
    @Streaming
    @GET("/v2/pet/findByStatus")
    suspend fun findPetsByStatusStreaming(
        
        
        @Query("status") status: kotlin.Array<kotlin.String>,
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Finds Pets by tags
     * Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
     *
     * @param tags Tags to filter by
     * @return kotlin.Array<Pet>
     */
    @GET("/v2/pet/findByTags")
    suspend fun findPetsByTags(
        
        
        @Query("tags") tags: kotlin.Array<kotlin.String>,
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<kotlin.Array<Pet>>

    /**
     * Finds Pets by tags
     * Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.
     *
     * @param tags Tags to filter by
     * @return kotlin.Array<Pet>
     */
    @Streaming
    @GET("/v2/pet/findByTags")
    suspend fun findPetsByTagsStreaming(
        
        
        @Query("tags") tags: kotlin.Array<kotlin.String>,
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Find pet by ID
     * Returns a single pet
     *
     * @param petId ID of pet to return/ example :: 789
     * @return Pet
     */
    @GET("/v2/pet/{petId}")
    suspend fun getPetById(
        
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<Pet>

    /**
     * Find pet by ID
     * Returns a single pet
     *
     * @param petId ID of pet to return/ example :: 789
     * @return Pet
     */
    @Streaming
    @GET("/v2/pet/{petId}")
    suspend fun getPetByIdStreaming(
        
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Update an existing pet
     * 
     *
     * @param body Pet object that needs to be added to the store
     * @return void
     */
    @PUT("/v2/pet")
    suspend fun updatePet(
        
        
        
        @Body body: Pet,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Update an existing pet
     * 
     *
     * @param body Pet object that needs to be added to the store
     * @return void
     */
    @Streaming
    @PUT("/v2/pet")
    suspend fun updatePetStreaming(
        
        
        
        @Body body: Pet,
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Updates a pet in the store with form data
     * 
     *
     * @param petId ID of pet that needs to be updated/ example :: 789
     * @param name Updated name of the pet/ example :: name_example
     * @param status Updated status of the pet/ example :: status_example
     * @return void
     */
    @POST("/v2/pet/{petId}")
    suspend fun updatePetWithForm(
        
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Updates a pet in the store with form data
     * 
     *
     * @param petId ID of pet that needs to be updated/ example :: 789
     * @param name Updated name of the pet/ example :: name_example
     * @param status Updated status of the pet/ example :: status_example
     * @return void
     */
    @Streaming
    @POST("/v2/pet/{petId}")
    suspend fun updatePetWithFormStreaming(
        
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * uploads an image
     * 
     *
     * @param petId ID of pet to update/ example :: 789
     * @param additionalMetadata Additional data to pass to server/ example :: additionalMetadata_example
     * @param file file to upload/ example :: /path/to/file.txt
     * @return ApiResponse
     */
    @POST("/v2/pet/{petId}/uploadImage")
    suspend fun uploadFile(
        
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ApiResponse>

    /**
     * uploads an image
     * 
     *
     * @param petId ID of pet to update/ example :: 789
     * @param additionalMetadata Additional data to pass to server/ example :: additionalMetadata_example
     * @param file file to upload/ example :: /path/to/file.txt
     * @return ApiResponse
     */
    @Streaming
    @POST("/v2/pet/{petId}/uploadImage")
    suspend fun uploadFileStreaming(
        
        @Path("petId") petId: kotlin.Long,
        
        
        @HeaderMap  customHeaders: Map<String, String> = emptyMap(),
        @QueryMap   customQueries: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    companion object {
        /**
         * Default PetApi Factory.
         */
        fun create(baseUrl: String, okHttpClient: OkHttpClient? = null, moshi: Moshi = InternalUtils.moshi, block: (builder: Retrofit.Builder)->Retrofit.Builder = { it }): PetApi {
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
            return  block(builder).build().create(PetApi::class.java)
        }
    }
}

