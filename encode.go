package jsons

import (
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/vuuvv/errors"
)

var Marshal = sonic.Marshal

func Stringify(value any) (string, error) {
	switch val := value.(type) {
	case string:
		return val, nil
	case []byte:
		return B2S(val), nil
	}
	bs, err := sonic.Marshal(value)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return B2S(bs), nil
}

func MustStringify(value any) string {
	ret, err := Stringify(value)
	if err != nil {
		panic(err)
	}
	return ret
}

func StringifyBytes(value any) ([]byte, error) {
	switch val := value.(type) {
	case string:
		return S2B(val), nil
	case []byte:
		return val, nil
	}
	bs, err := sonic.Marshal(value)
	return bs, errors.WithStack(err)
}

func MustStringifyBytes(value any) []byte {
	ret, err := StringifyBytes(value)
	if err != nil {
		panic(err)
	}
	return ret
}

func GetString(json string, path ...any) (string, error) {
	node, err := ast.NewSearcher(json).GetByPath(path...)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return node.String()
}

func GetBool(json string, path ...any) (bool, error) {
	node, err := ast.NewSearcher(json).GetByPath(path...)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return node.Bool()
}

func GetInt64(json string, path ...any) (int64, error) {
	node, err := ast.NewSearcher(json).GetByPath(path...)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return node.Int64()
}

func GetFloat64(json string, path ...any) (float64, error) {
	node, err := ast.NewSearcher(json).GetByPath(path...)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return node.Float64()
}
