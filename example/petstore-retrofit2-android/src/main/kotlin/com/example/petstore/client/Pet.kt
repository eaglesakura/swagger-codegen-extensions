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
data class Pet(
    /**
     * 
     */
    @Json(name = "id")
    val id: kotlin.Long? = null,
    /**
     * 
     */
    @Json(name = "category")
    val category: Category? = null,
    /**
     * 
     */
    @Json(name = "name")
    val name: kotlin.String,
    /**
     * 
     */
    @Json(name = "photoUrls")
    val photoUrls: kotlin.Array<kotlin.String>,
    /**
     * 
     */
    @Json(name = "tags")
    val tags: kotlin.Array<Tag>? = null,
    /**
     * pet status in the store
     */
    @Json(name = "status")
    val status: PetStatusEnum? = null,
) : android.os.Parcelable {

    constructor(origin: Pet): this(
            id = origin.id,
            category = origin.category,
            name = origin.name,
            photoUrls = origin.photoUrls,
            tags = origin.tags,
            status = origin.status,
    )

    constructor(json: String) : this(
            origin = requireNotNull(parse(json)) {
                "json parse failed<Pet>"
            }
    )

    constructor(parcel: android.os.Parcel) : this(
            json = requireNotNull(parcel.readString()) {
                "invalid Parcel json<Pet>"
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
        InternalUtils.moshi.adapter(Pet::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        @JvmStatic
        fun parse(json: String): Pet? {
            return InternalUtils.moshi.adapter(Pet::class.java).fromJson(json)
        }

        @JvmStatic
        val CREATOR: android.os.Parcelable.Creator<Pet> =
            object : android.os.Parcelable.Creator<Pet> {
                override fun createFromParcel(source: android.os.Parcel): Pet = Pet(source)
                override fun newArray(size: Int): Array<Pet> = arrayOf()
            }

    }
}


data class PetStatusEnum(private val value: String) {
    override fun toString(): String {
        return value
    }
    companion object {

        val available = PetStatusEnum("available")

        val pending = PetStatusEnum("pending")

        val sold = PetStatusEnum("sold")

        val MoshiTypeAdapter = object : JsonAdapter<PetStatusEnum>() {
            @Throws(IOException::class)
            override fun fromJson(reader: JsonReader): PetStatusEnum {
                return parse(reader.nextString())
            }

            @Throws(IOException::class)
            override fun toJson(writer: JsonWriter, value: PetStatusEnum?) {
                if(value == null) {
                    return
                }
                writer.value(value.toString())
            }

            override fun toString(): String {
                return "JsonAdapter(PetStatusEnum)"
            }
        }
        fun parse(value: String): PetStatusEnum {
            return when(value) {
                "" -> throw IllegalArgumentException("Invalid enum($value)")
                
                "available" -> available
                
                "pending" -> pending
                
                "sold" -> sold
                
                else -> PetStatusEnum(value)
            }
        }
    }
}

