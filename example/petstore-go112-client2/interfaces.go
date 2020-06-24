package petstore

import (
	"bytes"
	"io"
)

// link: https://github.com/eaglesakura/swagger-codegen-extensions

type swaggerModel interface {
	Json() string
}

func bodyToReader(body interface{}) io.Reader {
	switch body := body.(type) {
	case *io.Reader:
		return *body
	case io.Reader:
		return body
	case swaggerModel:
		return bytes.NewReader([]byte(body.Json()))
	}
	return nil
}
