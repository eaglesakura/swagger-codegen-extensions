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
    val complete: kotlin.Boolean? = null,
) : android.os.Parcelable {

    constructor(origin: Order): this(
            id = origin.id,
            petId = origin.petId,
            quantity = origin.quantity,
            shipDate = origin.shipDate,
            status = origin.status,
            complete = origin.complete,
    )

    constructor(json: String) : this(
            origin = requireNotNull(parse(json)) {
                "json parse failed<Order>"
            }
    )

    constructor(parcel: android.os.Parcel) : this(
            json = requireNotNull(parcel.readString()) {
                "invalid Parcel json<Order>"
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
        InternalUtils.moshi.adapter(Order::class.java).toJson(this)

    companion object {
        /**
         * Try parse json.
         */
        @JvmStatic
        fun parse(json: String): Order? {
            return InternalUtils.moshi.adapter(Order::class.java).fromJson(json)
        }

        @JvmStatic
        val CREATOR: android.os.Parcelable.Creator<Order> =
            object : android.os.Parcelable.Creator<Order> {
                override fun createFromParcel(source: android.os.Parcel): Order = Order(source)
                override fun newArray(size: Int): Array<Order> = arrayOf()
            }

    }
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

