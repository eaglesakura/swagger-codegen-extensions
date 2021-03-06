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
    val name: kotlin.String? = null,
)  {

    constructor(origin: Category): this(
            id = origin.id,
            name = origin.name,
    )

    constructor(json: String) : this(
            origin = requireNotNull(parse(json)) {
                "json parse failed<Category>"
            }
    )


    /**
     * Convert to Json.
     */
    fun toJson(): String =
        InternalUtils.moshi.adapter(Category::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        @JvmStatic
        fun parse(json: String): Category? {
            return InternalUtils.moshi.adapter(Category::class.java).fromJson(json)
        }


    }
}


