package com.eaglesakura.swagger2.languages

import com.eaglesakura.swagger2.extensions.normalize
import io.swagger.codegen.CliOption
import io.swagger.codegen.CodegenConstants
import io.swagger.codegen.CodegenType
import io.swagger.codegen.SupportingFile
import io.swagger.codegen.languages.AbstractKotlinCodegen
import io.swagger.models.Swagger
import io.swagger.models.parameters.BodyParameter
import java.io.File

class Retrofit2ClientCodegen : AbstractKotlinCodegen() {
    init {
        artifactId = "kotlin-retrofit2-client"
        packageName = "io.swagger.client"

        outputFolder = "generated-code" + File.separator + "kotlin-retrofit2-client"
        sourceFolder = "kotlin"
        modelTemplateFiles["model.mustache"] = ".kt"
        apiTemplateFiles["api.mustache"] = ".kt"
        embeddedTemplateDir = "kotlin-retrofit2-client".also { templateDir = it }
        apiPackage = packageName
        modelPackage = packageName

        cliOptions.add(CliOption(CodegenConstants.MODEL_PACKAGE, CodegenConstants.MODEL_PACKAGE_DESC))
        cliOptions.add(CliOption(CodegenConstants.API_PACKAGE, CodegenConstants.API_PACKAGE_DESC))
        cliOptions.add(CliOption(CodegenConstants.SOURCE_FOLDER, CodegenConstants.SOURCE_FOLDER_DESC))
    }

    override fun processOpts() {
        super.processOpts()
        supportingFiles.add(SupportingFile("internal_utils.mustache",
                ("$sourceFolder/$apiPackage").replace(".", "${File.separatorChar}"), "InternalUtils.kt"));
        supportingFiles.add(SupportingFile("api_enum_factory.mustache",
                ("$sourceFolder/$apiPackage").replace(".", "${File.separatorChar}"), "ApiEnumFactory.kt"));

        supportingFiles.add(SupportingFile("dependencies.gradle.mustache",
                sourceFolder.replace(".", "${File.separatorChar}"), "dependencies.gradle"));

    }

    override fun preprocessSwagger(swagger: Swagger) {
        super.preprocessSwagger(swagger)
        swagger.normalize()
        swagger.basePath = if (swagger.basePath?.endsWith("/") == true) {
            swagger.basePath!!.substring(0, swagger.basePath!!.length - 1)
        } else {
            swagger.basePath ?: ""
        }
    }

    override fun getTag(): CodegenType = CodegenType.CLIENT

    override fun getHelp(): String = "Generates a Retrofit2/Kotlin client."

    override fun getName(): String = "retrofit2"
}