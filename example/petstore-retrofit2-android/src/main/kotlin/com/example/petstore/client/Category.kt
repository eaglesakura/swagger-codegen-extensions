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
) : android.os.Parcelable {

    constructor(origin: Category): this(
            id = origin.id,
            name = origin.name,
    )

    constructor(json: String) : this(
            origin = requireNotNull(parse(json)) {
                "json parse failed<Category>"
            }
    )

    constructor(parcel: android.os.Parcel) : this(
            json = requireNotNull(parcel.readString()) {
                "invalid Parcel json<Category>"
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
        InternalUtils.moshi.adapter(Category::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        @JvmStatic
        fun parse(json: String): Category? {
            return InternalUtils.moshi.adapter(Category::class.java).fromJson(json)
        }

        @JvmStatic
        val CREATOR: android.os.Parcelable.Creator<Category> =
            object : android.os.Parcelable.Creator<Category> {
                override fun createFromParcel(source: android.os.Parcel): Category = Category(source)
                override fun newArray(size: Int): Array<Category> = arrayOf()
            }

    }
}


