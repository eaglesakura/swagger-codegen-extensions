package com.example.petstore.client

import java.math.* // ktlint-disable
import java.util.* // ktlint-disable
import java.io.* // ktlint-disable
import com.squareup.moshi.* // ktlint-disable

/**
 * 
 *
 * for Android model.
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
) : android.os.Parcelable {

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

    constructor(parcel: android.os.Parcel) : this(
            json = requireNotNull(parcel.readString()) {
                "invalid Parcel json<User>"
            }
    )

    override fun describeContents(): Int = 0

    override fun writeToParcel(dest: android.os.Parcel, flags: Int) {
        dest.writeString(toJson())
    }

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

        @JvmStatic
        val CREATOR: android.os.Parcelable.Creator<User> =
            object : android.os.Parcelable.Creator<User> {
                override fun createFromParcel(source: android.os.Parcel): User = User(source)
                override fun newArray(size: Int): Array<User> = arrayOf()
            }

    }
}


