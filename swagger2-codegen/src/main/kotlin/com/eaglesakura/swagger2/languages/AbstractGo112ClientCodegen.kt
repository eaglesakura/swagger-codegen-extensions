package com.eaglesakura.swagger2.languages

import com.eaglesakura.swagger2.extensions.normalize
import io.swagger.codegen.CliOption
import io.swagger.codegen.CodegenConstants
import io.swagger.codegen.languages.AbstractGoCodegen
import io.swagger.models.Swagger

abstract class AbstractGo112ClientCodegen : AbstractGoCodegen() {
    init {
        typeMapping["File"] = "io.Reader"
        typeMapping["file"] = "io.Reader"
        cliOptions.add(CliOption(CodegenConstants.PACKAGE_NAME, "Go package name (convention: lowercase).")
                .defaultValue("swagger"))
        cliOptions.add(CliOption(CodegenConstants.PACKAGE_VERSION, "Go package version.")
                .defaultValue("1.0.0"))
        cliOptions.add(CliOption(CodegenConstants.HIDE_GENERATION_TIMESTAMP, "hides the timestamp when files were generated")
                .defaultValue(true.toString()))
    }

    override fun preprocessSwagger(swagger: Swagger) {
        super.preprocessSwagger(swagger)
        swagger.normalize()
    }

    override fun toParamName(name: String?): String {
        return camelize(toVarName(name), false)
    }
}