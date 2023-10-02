package com.eaglesakura.swagger2.extensions

import io.swagger.models.Swagger
import io.swagger.models.parameters.BodyParameter
import io.swagger.models.properties.ObjectProperty
import io.swagger.models.properties.Property

private fun normalizePathTitle(s: String): String {
    return s.replace("{", "")
        .replace("}", "")
        .replace("/", "_")
        .replace("__", "_")
        .toLowerCase()
        .let {
            if (it.startsWith("_")) {
                it.substring(1)
            } else {
                it
            }
        }
}

private fun normalizeProperties(prefix: String, properties: Map<String, Property>) {
    properties
        .filterValues { property -> property is ObjectProperty }
        .forEach { (name, prop) ->
            val obj = prop as ObjectProperty
            obj.title = normalizePathTitle("${prefix}_$name")
            normalizeProperties(obj.title, obj.properties)
        }
}

/**
 * Update Swagger data.
 */
fun Swagger.normalize() {
    paths.forEach { (path, pathModel) ->
        pathModel.operationMap.forEach { (method, op) ->
            op.parameters
                .filterIsInstance(BodyParameter::class.java)
                .forEach { param ->
                    val schema = param.schema ?: return@forEach
                    schema.title = normalizePathTitle("${path}_${method}_${param.name}")

                    val properties = schema.properties ?: return@forEach
                    normalizeProperties(schema.title, properties)
                }
            op.responses
                .forEach { (status, resp) ->
                    val responseSchema = resp.responseSchema ?: return@forEach
                    responseSchema.title = normalizePathTitle("${path}_${method}_$status")

                    val properties = responseSchema.properties ?: return@forEach
                    normalizeProperties(responseSchema.title, properties)
                }
        }
    }
}
