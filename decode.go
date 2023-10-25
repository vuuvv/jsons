package jsons

import (
	"github.com/bytedance/sonic"
	"github.com/vuuvv/errors"
)

var Unmarshal = sonic.Unmarshal

func Parse[T any](json string) (*T, error) {
	return ParseBytes[T](S2B(json))
}

func MustParse[T any](json string) *T {
	return MustParseBytes[T](S2B(json))
}

func ParseBytes[T any](json []byte) (*T, error) {
	var ret T
	err := sonic.Unmarshal(json, &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ret, nil
}

func MustParseBytes[T any](json []byte) *T {
	val, err := ParseBytes[T](json)
	if err != nil {
		panic(err)
	}
	return val
}

func ParsePrimitive[T any](json string) (T, error) {
	return ParseBytesPrimitive[T](S2B(json))
}

func MustParsePrimitive[T any](json string) T {
	return MustParseBytesPrimitive[T]([]byte(json))
}

func ParseBytesPrimitive[T any](json []byte) (T, error) {
	var ret T
	err := sonic.Unmarshal(json, &ret)
	if err != nil {
		return ret, errors.WithStack(err)
	}
	return ret, nil
}

func MustParseBytesPrimitive[T any](json []byte) T {
	val, err := ParseBytesPrimitive[T](json)
	if err != nil {
		panic(err)
	}
	return val
}
