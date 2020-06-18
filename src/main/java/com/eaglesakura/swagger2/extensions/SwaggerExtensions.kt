package com.eaglesakura.swagger2.extensions

import io.swagger.models.Swagger
import io.swagger.models.parameters.BodyParameter

/**
 * Update Swagger data.
 */
fun Swagger.normalize() {
    paths.forEach { (path, pathModel) ->
        pathModel.operationMap.forEach { (method, op) ->
            op.parameters
                    .filterIsInstance(BodyParameter::class.java)
                    .filter { it.schema.properties.isNotEmpty() }
                    .forEach { param ->
                        param.schema.title = "${path}_${method}_${param.name}"
                                .replace("/", "_")
                                .toLowerCase()
                    }
            op.responses.forEach { (status, resp) ->
                if (resp.responseSchema?.properties?.isNotEmpty() == true) {
                    resp.responseSchema.title = "${path}_${method}_${status}"
                            .replace("/", "_")
                            .toLowerCase()
                }
            }
        }
    }
}