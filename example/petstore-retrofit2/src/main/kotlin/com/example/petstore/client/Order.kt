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
data class Order(

    /**
     * 
     */
    @Json(name = "id")
    val id: kotlin.Long? = null, 

    /**
     * 
     */
    @Json(name = "petId")
    val petId: kotlin.Long? = null, 

    /**
     * 
     */
    @Json(name = "quantity")
    val quantity: kotlin.Int? = null, 

    /**
     * 
     */
    @Json(name = "shipDate")
    val shipDate: java.time.LocalDateTime? = null, 

    /**
     * Order Status
     */
    @Json(name = "status")
    val status: OrderStatusEnum? = null, 

    /**
     * 
     */
    @Json(name = "complete")
    val complete: kotlin.Boolean? = null
) {
    companion object
}


data class OrderStatusEnum(private val value: String) {
    override fun toString(): String {
        return value
    }
    companion object {

        val placed = OrderStatusEnum("placed")

        val approved = OrderStatusEnum("approved")

        val delivered = OrderStatusEnum("delivered")

        val MoshiTypeAdapter = object : JsonAdapter<OrderStatusEnum>() {
            @Throws(IOException::class)
            override fun fromJson(reader: JsonReader): OrderStatusEnum {
                return parse(reader.nextString())
            }

            @Throws(IOException::class)
            override fun toJson(writer: JsonWriter, value: OrderStatusEnum?) {
                if(value == null) {
                    return
                }
                writer.value(value.toString())
            }

            override fun toString(): String {
                return "JsonAdapter(OrderStatusEnum)"
            }
        }
        fun parse(value: String): OrderStatusEnum {
            return when(value) {
                "" -> throw IllegalArgumentException("Invalid enum($value)")
                
                "placed" -> placed
                
                "approved" -> approved
                
                "delivered" -> delivered
                
                else -> OrderStatusEnum(value)
            }
        }
    }
}

