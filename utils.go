package jsons

import (
	"github.com/vuuvv/errors"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// B2S converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// S2B converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func S2B(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len

	return b
}

func Str2Int(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}
	if strings.Contains(str, ".") {
		n, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		return int64(n), nil
	} else {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, errors.WithStack(err)
		}
		return n, nil
	}
}
