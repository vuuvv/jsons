package jsons

import (
	"github.com/vuuvv/errors"
	"strconv"
	"strings"
	"time"
)

type Int struct {
	int
}

func (i *Int) ToInt() int {
	return i.int
}

func (i *Int) UnmarshalJSON(bs []byte) (err error) {
	str := B2S(bs)
	str = strings.TrimSpace(str)

	if str == "" {
		i.int = 0
	}

	var n int64
	// 非字符串格式,直接解析
	if str[0] != '"' {
		n, err = Str2Int(str)
		i.int = int(n)
		return
	}

	if str[len(str)-1] != '"' {
		return errors.Errorf("Int64格式错误: Need \" but %c", str[len(str)-1])
	}

	str = str[1 : len(str)-1]
	n, err = Str2Int(str)
	i.int = int(n)
	return
}

type Int64 struct {
	int64
}

func (i *Int64) ToInt64() int64 {
	return i.int64
}

func (i *Int64) MarshalJSON() ([]byte, error) {
	return S2B(strconv.FormatInt(i.int64, 10)), nil
}

func (i *Int64) UnmarshalJSON(bs []byte) (err error) {
	str := B2S(bs)
	str = strings.TrimSpace(str)

	if str == "" {
		i.int64 = 0
	}

	// 非字符串格式,直接解析
	if str[0] != '"' {
		i.int64, err = Str2Int(str)
		return
	}

	if str[len(str)-1] != '"' {
		return errors.Errorf("Int64格式错误: Need \" but %c", str[len(str)-1])
	}

	str = str[1 : len(str)-1]
	i.int64, err = Str2Int(str)
	return
}

type DateTime struct {
	time.Time
}

func (d *DateTime) ToTime() time.Time {
	return d.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	return S2B(d.Format("2006-01-02 15:04:05")), nil
}

func (d *DateTime) UnmarshalJSON(bs []byte) error {
	str := B2S(bs)
	str = strings.TrimSpace(str)
	if str[0] != '"' {
		return errors.Errorf("DateTime格式错误: Need \" but %c", str[0])
	}
	if str[len(str)-1] != '"' {
		return errors.Errorf("DateTime格式错误: Need \" but %c", str[len(str)-1])
	}
	str = str[1 : len(str)-1]

	if str == "" {
		return nil
	}

	var t time.Time
	var err error
	if len(str) == 10 {
		t, err = time.ParseInLocation("2006-01-02", str, time.Local)
	} else if len(str) == 19 {
		t, err = time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	} else {
		t, err = time.Parse("2006-01-02T15:04:05Z07:00", str)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	d.Time = t
	return nil
}
