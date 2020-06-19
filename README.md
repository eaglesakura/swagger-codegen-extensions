# Swagger Codegen Extensions

This library is [Swagger Codegen] custom language generator.

|Custom support language|Type|Target|Extra Library|Language Class|
|---|---|---|---|---|
|Kotlin|Client|Java VM, Android|Retrofit2|com.eaglesakura.swagger2.languages.Retrofit2ClientCodegen|
|Golang|Server|Go1.12|[swagger-go-core](https://github.com/eaglesakura/swagger-go-core)|com.eaglesakura.swagger2.languages.Go112ServerCodegen|
|Golang|Client|Go1.12|[swagger-go-core](https://github.com/eaglesakura/swagger-go-core)|com.eaglesakura.swagger2.languages.Go112ClientCodegen|
|Dart|Client|Dart2.0, Flutter|-|com.eaglesakura.swagger2.languages.Dart2ClientCodegen|

## Auto rename inline model.

This generator always rename inline-model.

Original swagger-codegen generate a bad name to inline-model(InlineResponse200, InlineResponse2001, InlineResponse2002...).

This library rename it to better.(PathToFooGet200, PathToBarGet200...)

# How to use / Swagger2.0

You can generate source codes by Gradle plugin.

https://plugins.gradle.org/plugin/org.hidetake.swagger.generator

```gradle
// see) example/build.gradle

plugins {
    id("org.hidetake.swagger.generator") version "2.18.2"
}

repositories {
    mavenLocal()
    google()
    jcenter()
    mavenCentral()
    maven { url "https://dl.bintray.com/eaglesakura/maven/" }
}

dependencies {
    swaggerCodegen 'com.eaglesakura.swagger.swagger2-codegen:swagger2-codegen:1.0.build-24'
}

swaggerSources {
    petstore_retrofit2 {
        inputFile = file("petstore.yaml")
        code {
            language = "com.eaglesakura.swagger2.languages.Retrofit2ClientCodegen"
            configFile = file("petstore-config-retrofit2.json")
            outputDir = file("petstore-retrofit2")
//            templateDir = file('path/to/custom/template/directory')
        }
    }
}

```

# License

```
The MIT License (MIT)

Copyright (c) 2020 @eaglesakura

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```