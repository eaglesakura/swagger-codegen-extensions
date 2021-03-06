package petstore

import (
	"bytes"
	"golang.org/x/xerrors"
	"io"
	"net/http"
	"strconv"
	"time"
)

//noinspection GoUnusedGlobalVariable
var GeneratedDate, _ = time.Parse(time.RFC3339, "2020-06-26T14:39:29.935+09:00")

// link: https://github.com/eaglesakura/swagger-codegen-extensions

type SwaggerResponse interface {
    Write(writer http.ResponseWriter, request *http.Request)
}

type swaggerValidatable interface {
    Valid() error
}

type swaggerModel interface {
	Json() []byte
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

func convertFloatValue(value string, result interface{}) error {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return xerrors.Errorf("ParseFloat(%v) failed: %w", value, err)
	}

	switch ptr := result.(type) {
	case **float32:
		v := float32(f)
		*ptr = &v
	case **float64:
		*ptr = &f
	}
	return nil
}

func convertIntValue(value string, result interface{}) error {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return xerrors.Errorf("ParseInt(%v) failed: %w", value, err)
	}

	switch ptr := result.(type) {
	case **int8:
		v := int8(i)
		*ptr = &v
	case **uint8:
		v := uint8(i)
		*ptr = &v
	case **int16:
		v := int16(i)
		*ptr = &v
	case **uint16:
		v := uint16(i)
		*ptr = &v
	case **int32:
		v := int32(i)
		*ptr = &v
	case **uint32:
		v := uint32(i)
		*ptr = &v
	case **int64:
		*ptr = &i
	case **uint64:
		v := uint64(i)
		*ptr = &v
	case **int:
		v := int(i)
		*ptr = &v
	case **uint:
		v := uint(i)
		*ptr = &v
	}
	return nil
}

func convertValue(value string, result interface{}) error {
	switch ptr := result.(type) {
	case **string:
		*ptr = &value
	case **int8, **uint8, **int16, **uint16, **int32, **uint32, **int64, **uint64, **int, **uint:
		return convertIntValue(value, result)
	case **float32, **float64:
		return convertFloatValue(value, result)
	case **bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return xerrors.Errorf("ParseBool(%v) failed: %w", value, err)
		}
		*ptr = &b
	default:
		return xerrors.Errorf("unsupported convert type(%v)", ptr)
	}
	return nil
}

func parsePath(values map[string]string, key string, result interface{}) error {
	value, ok := values[key]
	if !ok {
		return xerrors.Errorf("invalid path[%v] parameter", key)
	}
	return convertValue(value, result)
}

func parseValues(values map[string][]string, key string, result interface{}) error {
	value, ok := values[key]
	if !ok || len(value) != 1 {
		return xerrors.Errorf("invalid values[%v] parameter", key)
	}
	return convertValue(value[0], result)
}

func validationValue(ref interface{}) error {
    validatable, ok := ref.(swaggerValidatable)
    if !ok || validatable == nil {
        return nil
    }

    err := validatable.Valid()
    if err != nil {
        return xerrors.Errorf(": %w", err)
    }
    return nil
}