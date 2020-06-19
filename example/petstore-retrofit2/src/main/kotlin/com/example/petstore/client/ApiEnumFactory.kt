package com.example.petstore.client

import com.example.petstore.client.* // ktlint-disable
import java.math.* // ktlint-disable
import java.util.* // ktlint-disable
import java.io.* // ktlint-disable
import com.squareup.moshi.* // ktlint-disable
import retrofit2.* // ktlint-disable
import retrofit2.http.* // ktlint-disable
import java.lang.reflect.Type

object ApiEnumFactory : JsonAdapter.Factory {
    override fun create(type: Type, annotations: Set<out Annotation>, moshi: Moshi): JsonAdapter<*>? {
        return when(type) {
            
            
                OrderStatusEnum::class.java -> OrderStatusEnum.MoshiTypeAdapter
            
                PetStatusEnum::class.java -> PetStatusEnum.MoshiTypeAdapter
            
            else -> null
        }
    }
}