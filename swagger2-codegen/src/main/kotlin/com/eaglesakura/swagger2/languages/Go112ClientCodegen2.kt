package com.eaglesakura.swagger2.languages

import io.swagger.codegen.CodegenType
import io.swagger.codegen.SupportingFile
import java.io.File

class Go112ClientCodegen2 : AbstractGo112ClientCodegen() {
    init {
        outputFolder = "generated-code" + File.separator + "go112-client2"
        embeddedTemplateDir = "go112-client2".also { templateDir = it }
        supportingFiles.add(SupportingFile("interfaces.mustache", "interfaces.go"))
    }

    override fun getTag(): CodegenType = CodegenType.CLIENT

    override fun getHelp(): String = "Generates a Go 1.12 client / Ver2."

    override fun getName(): String = "go112-client2"
}