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
data class Tag(
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

    constructor(origin: Tag): this(
            id = origin.id,
            name = origin.name,
    )

    internal constructor(parcel: android.os.Parcel) : this(
            json = requireNotNull(parse(parcel.readString() ?: "")) {
                "invalid Parcel json<Tag>"
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
        InternalUtils.moshi.adapter(Tag::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        @JvmStatic
        fun parse(json: String): Tag? {
            return InternalUtils.moshi.adapter(Tag::class.java).fromJson(json)
        }

        @JvmStatic
        internal val CREATOR: android.os.Parcelable.Creator<Tag> =
            object : android.os.Parcelable.Creator<Tag> {
                override fun createFromParcel(source: android.os.Parcel): Tag = Tag(source)
                override fun newArray(size: Int): Array<Tag> = arrayOf()
            }

    }
}


