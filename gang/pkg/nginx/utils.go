package nginx

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//AsSha256 hash for struct
func AsSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

//CovertToString convert interface to string
func CovertToString(value interface{}) string {
	if value == nil {
		return ""
	}
	var val string
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.String {
		val = v.String()
	} else if v.Type() == stringPtrType {
		val = v.Elem().String()
	} else if v.Kind() == reflect.Bool {
		if v.Bool() {
			val = "on"
		} else {
			val = "off"
		}
	} else if v.Kind() == reflect.Int {
		val = strconv.FormatInt(v.Int(), 10)
	} else if v.Kind() == reflect.Float64 {
		val = fmt.Sprintf("%g", v.Float())
	} else if v.Kind() == reflect.Slice {
		val = ""
		len := v.Len()
		for i := 0; i < len; i++ {
			if i == (len - 1) {
				val = val + v.Index(i).String()
			} else {
				val = val + v.Index(i).String() + " "
			}
		}
	}

	return val
}

//BuildDirective build directive with name , value
func BuildDirective(name string, value interface{}) Directive {
	t, ok := value.(*Block)
	if ok {
		return (Directive)(t)
	}

	return NewKeyValueOption(name, value)
}

//MarshalDirective marshal struct to directives such KeyValueOption
func MarshalDirective(i interface{}) ([]Directive, error) {
	v := getValue(reflect.ValueOf(i))
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type %s is not supported", t.Kind())
	}

	var directives []Directive
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// skip unexported fields. from godoc:
		// PkgPath is the package path that qualifies a lower case (unexported)
		// field name. It is empty for upper case (exported) field names.
		if f.PkgPath != "" {
			continue
		}
		fv := getValue(v.Field(i))
		if fv.IsValid() {

		}
		key, omit := readTag(f) //omit 忽略
		var d Directive
		switch fv.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice, reflect.Ptr, reflect.UnsafePointer:
			if fv.IsNil() {
				if !omit {
					d = BuildDirective(key, nil)
				} else {
					continue
				}
			} else {
				d = BuildDirective(key, fv.Interface())
			}
		default:
			d = BuildDirective(key, fv.Interface())
		}

		directives = append(directives, d)
	}

	return directives, nil
}

//getValue 判断是否是Ptr如果是就调用Elm()
func getValue(value reflect.Value) reflect.Value {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return value
		}
		return getValue(value.Elem())
	default:
		return value
	}
}

// read tag like `kv:"email,omitempty"`
func readTag(f reflect.StructField) (string, bool) {
	val, ok := f.Tag.Lookup("kv")
	if !ok {
		return f.Name, false
	}
	opts := strings.Split(val, ",")
	omit := false
	if len(opts) == 2 {
		omit = opts[1] == "omitempty"
	}
	return opts[0], omit
}
