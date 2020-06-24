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
 * StoreApi Client.
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
 *      val exampleApi: StoreApi = StoreApi.create("https://example.com")
 *      exampleApi.fooApi()
 * }
 *
 * @link https://github.com/eaglesakura/swagger-codegen-extensions
 */
interface StoreApi {
    /**
     * Delete purchase order by ID
     * For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
     *
     * @param orderId ID of the order that needs to be deleted/ example :: 789
     * @return void
     */
    @DELETE("/v2/store/order/{orderId}")
    suspend fun deleteOrder(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Path("orderId")
            orderId: kotlin.Long
    ): Response<ResponseBody>

    /**
     * Delete purchase order by ID
     * For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
     *
     * @param orderId ID of the order that needs to be deleted/ example :: 789
     * @return void
     */
    @DELETE("/v2/store/order/{orderId}")
    suspend fun deleteOrderBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Path("orderId")
            orderId: kotlin.Long
    ): Response<ResponseBody>

    /**
     * Delete purchase order by ID
     * For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors
     *
     * @param orderId ID of the order that needs to be deleted/ example :: 789
     * @return void
     */
    @Streaming
    @DELETE("/v2/store/order/{orderId}")
    suspend fun deleteOrderStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Path("orderId")
            orderId: kotlin.Long
    ): Response<ResponseBody>
    /**
     * Returns pet inventories by status
     * Returns a map of status codes to quantities
     *
     * @return kotlin.collections.Map<kotlin.String, kotlin.Int>
     */
    @GET("/v2/store/inventory")
    suspend fun getInventory(
            @HeaderMap customHeaders: Map<String, String> = emptyMap()
    ): Response<kotlin.collections.Map<kotlin.String, kotlin.Int>>

    /**
     * Returns pet inventories by status
     * Returns a map of status codes to quantities
     *
     * @return kotlin.collections.Map<kotlin.String, kotlin.Int>
     */
    @GET("/v2/store/inventory")
    suspend fun getInventoryBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap()
    ): Response<ResponseBody>

    /**
     * Returns pet inventories by status
     * Returns a map of status codes to quantities
     *
     * @return kotlin.collections.Map<kotlin.String, kotlin.Int>
     */
    @Streaming
    @GET("/v2/store/inventory")
    suspend fun getInventoryStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap()
    ): Response<ResponseBody>
    /**
     * Find purchase order by ID
     * For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions
     *
     * @param orderId ID of pet that needs to be fetched/ example :: 789
     * @return Order
     */
    @GET("/v2/store/order/{orderId}")
    suspend fun getOrderById(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Path("orderId")
            orderId: kotlin.Long
    ): Response<Order>

    /**
     * Find purchase order by ID
     * For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions
     *
     * @param orderId ID of pet that needs to be fetched/ example :: 789
     * @return Order
     */
    @GET("/v2/store/order/{orderId}")
    suspend fun getOrderByIdBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Path("orderId")
            orderId: kotlin.Long
    ): Response<ResponseBody>

    /**
     * Find purchase order by ID
     * For valid response try integer IDs with value &gt;&#x3D; 1 and &lt;&#x3D; 10.         Other values will generated exceptions
     *
     * @param orderId ID of pet that needs to be fetched/ example :: 789
     * @return Order
     */
    @Streaming
    @GET("/v2/store/order/{orderId}")
    suspend fun getOrderByIdStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Path("orderId")
            orderId: kotlin.Long
    ): Response<ResponseBody>
    /**
     * Place an order for a pet
     * 
     *
     * @param body order placed for purchasing the pet
     * @return Order
     */
    @POST("/v2/store/order")
    suspend fun placeOrder(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: Order
    ): Response<Order>

    /**
     * Place an order for a pet
     * 
     *
     * @param body order placed for purchasing the pet
     * @return Order
     */
    @POST("/v2/store/order")
    suspend fun placeOrderBody(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: Order
    ): Response<ResponseBody>

    /**
     * Place an order for a pet
     * 
     *
     * @param body order placed for purchasing the pet
     * @return Order
     */
    @Streaming
    @POST("/v2/store/order")
    suspend fun placeOrderStreaming(
            @HeaderMap customHeaders: Map<String, String> = emptyMap(),
            @Body 
            body: Order
    ): Response<ResponseBody>
    companion object {
        /**
         * Default StoreApi Factory.
         */
        fun create(baseUrl: String, okHttpClient: OkHttpClient? = null, moshi: Moshi = InternalUtils.moshi, block: (builder: Retrofit.Builder)->Retrofit.Builder = { it }): StoreApi {
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
            return  block(builder).build().create(StoreApi::class.java)
        }
    }
}

