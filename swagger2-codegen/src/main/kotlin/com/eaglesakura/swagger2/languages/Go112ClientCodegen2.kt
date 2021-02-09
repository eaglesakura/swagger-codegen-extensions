package com.eaglesakura.swagger2.languages

import io.swagger.codegen.CodegenType
import io.swagger.codegen.SupportingFile
import java.io.File

class Go112ClientCodegen2 : AbstractGo112ClientCodegen() {
    init {
        outputFolder = "generated-code" + File.separator + "go112-client2"
        embeddedTemplateDir = "go112-v2".also { templateDir = it }
        modelTemplateFiles["model.mustache"] = ".go"
        apiTemplateFiles["api-model.mustache"] = "_model.go"
        apiTemplateFiles["api-client.mustache"] = "_client.go"
        supportingFiles.add(SupportingFile("go-client.mod.mustache", "go.mod"))
        supportingFiles.add(SupportingFile("utils.mustache", "utils.go"))
    }

    override fun getTag(): CodegenType = CodegenType.CLIENT

    override fun getHelp(): String = "Generates a Go 1.12 client / Ver2."

    override fun getName(): String = "go112-client2"
}
