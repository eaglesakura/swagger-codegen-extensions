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
data class User(
    /**
     * 
     */
    @Json(name = "id")
    val id: kotlin.Long? = null,
    /**
     * 
     */
    @Json(name = "username")
    val username: kotlin.String? = null,
    /**
     * 
     */
    @Json(name = "firstName")
    val firstName: kotlin.String? = null,
    /**
     * 
     */
    @Json(name = "lastName")
    val lastName: kotlin.String? = null,
    /**
     * 
     */
    @Json(name = "email")
    val email: kotlin.String? = null,
    /**
     * 
     */
    @Json(name = "password")
    val password: kotlin.String? = null,
    /**
     * 
     */
    @Json(name = "phone")
    val phone: kotlin.String? = null,
    /**
     * User Status
     */
    @Json(name = "userStatus")
    val userStatus: kotlin.Int? = null,
)  {

    constructor(origin: User): this(
            id = origin.id,
            username = origin.username,
            firstName = origin.firstName,
            lastName = origin.lastName,
            email = origin.email,
            password = origin.password,
            phone = origin.phone,
            userStatus = origin.userStatus,
    )

    constructor(json: String) : this(
            origin = requireNotNull(parse(json)) {
                "json parse failed<User>"
            }
    )


    /**
     * Convert to Json.
     */
    fun toJson(): String =
        InternalUtils.moshi.adapter(User::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        @JvmStatic
        fun parse(json: String): User? {
            return InternalUtils.moshi.adapter(User::class.java).fromJson(json)
        }


    }
}


