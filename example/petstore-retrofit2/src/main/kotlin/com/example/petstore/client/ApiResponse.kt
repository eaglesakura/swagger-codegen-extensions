package com.example.petstore.client

import java.math.* // ktlint-disable
import java.util.* // ktlint-disable
import java.io.* // ktlint-disable
import com.squareup.moshi.* // ktlint-disable

/**
 * 
 *
 * 
 * @link https://github.com/eaglesakura/swagger-codegen-extensions
 **/
@Suppress("unused")
data class ApiResponse(

    /**
     * 
     */
    @Json(name = "code")
    val code: kotlin.Int? = null, 

    /**
     * 
     */
    @Json(name = "type")
    val type: kotlin.String? = null, 

    /**
     * 
     */
    @Json(name = "message")
    val message: kotlin.String? = null
) {

    /**
     * Convert to Json.
     */
    fun toJson(): String =
        InternalUtils.moshi.adapter(ApiResponse::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        fun parse(json: String): ApiResponse? {
            return InternalUtils.moshi.adapter(ApiResponse::class.java).fromJson(json)
        }
    }
}


