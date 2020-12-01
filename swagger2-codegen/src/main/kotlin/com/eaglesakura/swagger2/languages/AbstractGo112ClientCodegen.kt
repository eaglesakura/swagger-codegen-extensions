package com.eaglesakura.swagger2.languages

import com.eaglesakura.swagger2.extensions.normalize
import io.swagger.codegen.CliOption
import io.swagger.codegen.CodegenConstants
import io.swagger.codegen.CodegenProperty
import io.swagger.codegen.languages.AbstractGoCodegen
import io.swagger.models.Swagger

abstract class AbstractGo112ClientCodegen : AbstractGoCodegen() {

    val GO_MODULE_NAME = "goModuleName"
    val GO_MODULE_NAME_DESC = "module name for generated go module code"

    @JvmField
    protected var goModuleName: String = "swagger"

    init {
        typeMapping["File"] = "io.Reader"
        typeMapping["file"] = "io.Reader"
        cliOptions.add(CliOption(CodegenConstants.PACKAGE_NAME, "Go package name (convention: lowercase).")
                .defaultValue("swagger"))
        cliOptions.add(CliOption(GO_MODULE_NAME, GO_MODULE_NAME_DESC)
                .defaultValue("example.com/swagger"))
        cliOptions.add(CliOption(CodegenConstants.PACKAGE_VERSION, "Go package version.")
                .defaultValue("1.0.0"))
        cliOptions.add(CliOption(CodegenConstants.HIDE_GENERATION_TIMESTAMP, "hides the timestamp when files were generated")
                .defaultValue(true.toString()))
    }

    override fun cliOptions(): MutableList<CliOption> {
        return super.cliOptions()
    }

    override fun processOpts() {
        super.processOpts()

        if (additionalProperties.containsKey(GO_MODULE_NAME)) {
            goModuleName = additionalProperties[GO_MODULE_NAME]?.toString() ?: "swagger"
        }

        if (!additionalProperties.containsKey(CodegenConstants.PACKAGE_NAME)) {
            val packageName = goModuleName.split("/").last()
            LOGGER.info("Go package name -> $packageName")
            additionalProperties[CodegenConstants.PACKAGE_NAME] = packageName
        }
    }

    override fun preprocessSwagger(swagger: Swagger) {
        super.preprocessSwagger(swagger)
        swagger.normalize()
    }

    override fun toParamName(name: String?): String {
        return camelize(toVarName(name), false)
    }

    override fun toEnumName(property: CodegenProperty?): String {
        val result = super.toEnumName(property).let {
            it[0] + it.substring(1).toLowerCase()
        }
        return result
    }
}