package father

import (
	"strings"
	"regexp"
	"reflect"
	"fmt"
	"strconv"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

func _if(cond bool, v1, v2 interface{}) interface{} {
	if cond {
		return v1
	}

	return v2
}

func sliceFilter(items []string) []string {
	result := []string{}
	for _, text := range items {
		if strings.Trim(text, " ") != "" {
			result = append(result, text)
		}
	}

	return result
}

func sliceIn(elt, slice interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(slice))

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elt) {
			return true
		}
	}

	return false
}

func bytesClone(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)

	return c
}

func toSnakeCase(str string, toLower bool) string {
	snake := matchAllCap.ReplaceAllString(matchFirstCap.ReplaceAllString(str, "${1}_${2}"), "${1}_${2}")

	if toLower {
		return strings.ToLower(snake)
	}

	return snake
}

// 首字母小写
func lcFirst(n string) string {
	return strings.ToLower(n[:1]) + n[1:]
}

func convertAssign(dest, src interface{}) error {
	switch s := src.(type) {
	case string:
		switch d := dest.(type) {
		case *string:
			*d = s
		case *[]byte:
			*d = []byte(s)
		default:
			return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", src, dest)
		}

		return nil
	case []byte:
		switch d := dest.(type) {
		case *string:
			*d = string(s)
		case *interface{}:
			*d = bytesClone(s)
		case *[]byte:
			*d = bytesClone(s)
		default:
			return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", src, dest)
		}

		return nil
	case nil:
		switch d := dest.(type) {
		case *interface{}:
			*d = nil
		case *[]byte:
			*d = nil
		default:
			return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", src, dest)
		}

		return nil
	}

	switch d := dest.(type) {
	case *string:
		*d = src.(string)
	case *[]byte:
		*d = src.([]byte)
	case *bool:
		*d = src.(bool)
	default:
		dv := reflect.Indirect(reflect.ValueOf(dest))

		switch dv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var s string

			// 解决 php 传过来数据类型问题
			v := reflect.ValueOf(src)
			switch v.Kind() {
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				s = strconv.FormatInt(int64(v.Uint()), 10)
			default:
				s = strconv.FormatInt(v.Int(), 10)
			}

			i64, err := strconv.ParseInt(s, 10, dv.Type().Bits())
			if err != nil {
				return strconvErr(src, s, dv.Kind(), err)
			}
			dv.SetInt(i64)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			s := strconv.FormatUint(reflect.ValueOf(src).Uint(), 10)
			u64, err := strconv.ParseUint(s, 10, dv.Type().Bits())
			if err != nil {
				return strconvErr(src, s, dv.Kind(), err)
			}
			dv.SetUint(u64)
		case reflect.Float32, reflect.Float64:
			var s string
			rv := reflect.ValueOf(src)
			switch rv.Kind() {
			case reflect.Float64:
				s = strconv.FormatFloat(rv.Float(), 'g', -1, 64)
			case reflect.Float32:
				s = strconv.FormatFloat(rv.Float(), 'g', -1, 32)
			}

			f64, err := strconv.ParseFloat(s, dv.Type().Bits())
			if err != nil {
				return strconvErr(src, s, dv.Kind(), err)
			}
			dv.SetFloat(f64)
		default:
			return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", src, dest)
		}
	}

	return nil
}

func strconvErr(src interface{}, s string, kind reflect.Kind, err error) error {
	if ne, ok := err.(*strconv.NumError); ok {
		err = ne.Err
	}

	return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", src, s, kind, err)
}
