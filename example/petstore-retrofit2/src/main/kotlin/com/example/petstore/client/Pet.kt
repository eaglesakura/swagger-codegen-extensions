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
    val status: PetStatusEnum? = null
) {

    /**
     * Convert to Json.
     */
    fun toJson(): String =
        InternalUtils.moshi.adapter(Pet::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        fun parse(json: String): Pet? {
            return InternalUtils.moshi.adapter(Pet::class.java).fromJson(json)
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

