package com.eaglesakura.swagger2.languages

import com.eaglesakura.swagger2.extensions.normalize
import io.swagger.codegen.CodegenType
import io.swagger.codegen.languages.DartClientCodegen
import io.swagger.models.Swagger

class Dart2ClientCodegen : DartClientCodegen() {

    init {
        embeddedTemplateDir = "dart2-client"
    }

    override fun preprocessSwagger(swagger: Swagger) {
        super.preprocessSwagger(swagger)
        swagger.normalize()
    }

    override fun getTag(): CodegenType = CodegenType.CLIENT

    override fun getHelp(): String = "Generates a Dart2 client."

    override fun getName(): String = "dart2-client"
}