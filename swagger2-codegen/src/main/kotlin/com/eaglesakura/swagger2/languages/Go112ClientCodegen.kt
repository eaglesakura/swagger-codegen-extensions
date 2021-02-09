package com.eaglesakura.swagger2.languages

import io.swagger.codegen.CodegenType
import io.swagger.codegen.SupportingFile
import java.io.File

@Deprecated("Replace to Go112ClientCodegen2")
class Go112ClientCodegen : AbstractGo112ClientCodegen() {
    init {
        outputFolder = "generated-code" + File.separator + "go112-client"
        embeddedTemplateDir = "go112-client".also { templateDir = it }
        modelTemplateFiles["model.mustache"] = ".go"
        apiTemplateFiles["api.mustache"] = ".go"
        supportingFiles.add(SupportingFile("go.mod.mustache", "go.mod"))
    }

    override fun getTag(): CodegenType = CodegenType.CLIENT

    override fun getHelp(): String = "Generates a Go 1.12 client."

    override fun getName(): String = "go112-client"
}
