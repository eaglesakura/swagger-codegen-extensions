package com.eaglesakura.swagger2.languages

import com.eaglesakura.swagger2.extensions.normalize
import io.swagger.codegen.CliOption
import io.swagger.codegen.CodegenConstants
import io.swagger.codegen.CodegenType
import io.swagger.codegen.SupportingFile
import io.swagger.codegen.languages.AbstractGoCodegen
import io.swagger.models.Swagger
import java.io.File

class Go112ClientCodegen : AbstractGoCodegen() {
    init {
        outputFolder = "generated-code" + File.separator + "go112-client"
        modelTemplateFiles["model.mustache"] = ".go"
        apiTemplateFiles["api.mustache"] = ".go"
        embeddedTemplateDir = "go112-client".also { templateDir = it }
        typeMapping["File"] = "io.Reader"
        typeMapping["file"] = "io.Reader"
        cliOptions.add(CliOption(CodegenConstants.PACKAGE_NAME, "Go package name (convention: lowercase).")
                .defaultValue("swagger"))
        cliOptions.add(CliOption(CodegenConstants.PACKAGE_VERSION, "Go package version.")
                .defaultValue("1.0.0"))
        cliOptions.add(CliOption(CodegenConstants.HIDE_GENERATION_TIMESTAMP, "hides the timestamp when files were generated")
                .defaultValue(true.toString()))
    }

    override fun processOpts() {
        super.processOpts()
        supportingFiles.add(SupportingFile("go.mod.mustache", "go.mod"))
    }

    override fun preprocessSwagger(swagger: Swagger) {
        super.preprocessSwagger(swagger)
        swagger.normalize()
    }

    override fun getTag(): CodegenType = CodegenType.CLIENT

    override fun getHelp(): String = "Generates a Go 1.12 client."

    override fun getName(): String = "go112-client"
}