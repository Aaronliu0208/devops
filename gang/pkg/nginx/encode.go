package nginx

import (
	"fmt"
	"reflect"
	"strings"
)

//Marshal marshal struct to directives such KeyValueOption
func Marshal(i interface{}) ([]Directive, error) {
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
		case reflect.Func, reflect.Chan:
			continue
		case reflect.Interface, reflect.Struct:
			b := NewBlock(key)
			dd, err := Marshal(fv.Interface())
			if err != nil {
				continue
			}
			for _, sd := range dd {
				b.AddDirective(sd)
			}
			d = b
		case reflect.Ptr, reflect.UnsafePointer:
			if fv.IsNil() {
				if !omit {
					d = BuildDirective(key, nil)
				} else {
					continue
				}
			} else {
				d = BuildDirective(key, fv.Interface())
			}
		case reflect.Slice:
			if fv.IsNil() {
				if !omit {
					d = BuildDirective(key, nil)
				} else {
					continue
				}
			} else {
				b := NewBlock(key)
				for i := 0; i < fv.Len(); i++ {
					ifv := fv.Index(i)
					dd, err := Marshal(ifv.Interface())
					if err != nil {
						continue
					}
					for _, sd := range dd {
						b.AddDirective(sd)
					}
				}
				d = b
			}

		case reflect.Map:

		default:
			d = BuildDirective(key, fv.Interface())
		}

		directives = append(directives, d)
	}

	return directives, nil
}

//getValue 判断是否是Ptr如果是就调用Elm()
func getValue(value reflect.Value) reflect.Value {
	if !value.IsValid() {
		return reflect.ValueOf(nil)
	}
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
