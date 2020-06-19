package com.example.petstore.client

import java.math.* // ktlint-disable
import java.util.* // ktlint-disable
import java.io.* // ktlint-disable
import com.squareup.moshi.* // ktlint-disable

/**
 * 
 *
 * @link https://github.com/eaglesakura/swagger-codegen-extensions
 **/
data class Category(

    /**
     * 
     */
    @Json(name = "id")
    val id: kotlin.Long? = null, 

    /**
     * 
     */
    @Json(name = "name")
    val name: kotlin.String? = null
) {
    companion object
}


