import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

buildscript {
    extra["kotlin_version"] = "1.3.72"
    repositories {
        google()
        jcenter()
        mavenCentral()
    }
    dependencies {
        classpath("org.jetbrains.dokka:dokka-gradle-plugin:0.9.18") // kotlin-docs
        classpath("com.jfrog.bintray.gradle:gradle-bintray-plugin:1.8.4")
    }
}

plugins {
    kotlin("jvm") version "${extra["kotlin_version"]}"
}

repositories {
    mavenLocal()
    google()
    jcenter()
    mavenCentral()
}

dependencies {
    /**
     * Kotlin support
     */
    api("org.jetbrains.kotlin:kotlin-stdlib-jdk7:${extra["kotlin_version"]}")
    api("io.swagger:swagger-codegen:2.4.14")
    testImplementation("org.jetbrains.kotlin:kotlin-test")
    testImplementation("org.jetbrains.kotlin:kotlin-test-junit")
}

val compileKotlin: KotlinCompile by tasks
val compileTestKotlin: KotlinCompile by tasks

compileKotlin.kotlinOptions {
    jvmTarget = "1.8"
}
compileTestKotlin.kotlinOptions {
    jvmTarget = "1.8"
}
