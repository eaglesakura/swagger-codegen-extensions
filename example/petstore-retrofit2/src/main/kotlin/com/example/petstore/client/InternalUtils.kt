package com.example.petstore.client

import com.example.petstore.client.* // ktlint-disable
import java.math.* // ktlint-disable
import java.util.* // ktlint-disable
import java.io.* // ktlint-disable
import com.squareup.moshi.kotlin.reflect.KotlinJsonAdapterFactory
import com.squareup.moshi.* // ktlint-disable
import retrofit2.* // ktlint-disable
import retrofit2.http.* // ktlint-disable

internal object InternalUtils {
    val moshi = Moshi.Builder()
        .add(ApiEnumFactory)
        .add(KotlinJsonAdapterFactory())
        .build()
}