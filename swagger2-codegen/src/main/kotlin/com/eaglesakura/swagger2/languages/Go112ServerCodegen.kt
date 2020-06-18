package com.eaglesakura.swagger2.languages

import com.eaglesakura.swagger2.extensions.normalize
import io.swagger.codegen.CliOption
import io.swagger.codegen.CodegenConstants
import io.swagger.codegen.CodegenType
import io.swagger.codegen.languages.AbstractGoCodegen
import io.swagger.models.Swagger
import java.io.File

class Go112ServerCodegen : AbstractGoCodegen() {
    init {
        outputFolder = "generated-code" + File.separator + "go112-server"
        modelTemplateFiles["model.mustache"] = ".go"
        apiTemplateFiles["api.mustache"] = ".go"
        embeddedTemplateDir = "go112-server".also { templateDir = it }
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

    override fun getTag(): CodegenType = CodegenType.SERVER

    override fun getHelp(): String = "Generates a Go 1.12 server."

    override fun getName(): String = "go112-server"
}