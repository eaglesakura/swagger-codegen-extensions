package com.eaglesakura.swagger2.languages

import io.swagger.codegen.CodegenType
import io.swagger.codegen.SupportingFile
import java.io.File

class Go112ServerCodegen : AbstractGo112ClientCodegen() {
    init {
        outputFolder = "generated-code" + File.separator + "go112-server"
        embeddedTemplateDir = "go112-server".also { templateDir = it }
        modelTemplateFiles["model.mustache"] = ".go"
        apiTemplateFiles["api.mustache"] = ".go"
        supportingFiles.add(SupportingFile("go.mod.mustache", "go.mod"))
    }

    override fun getTag(): CodegenType = CodegenType.SERVER

    override fun getHelp(): String = "Generates a Go 1.12 server."

    override fun getName(): String = "go112-server"
}