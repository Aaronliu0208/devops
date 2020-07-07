package nginx

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
)

//AsSha256 hash for struct
func AsSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

//CovertToString convert interface to string
func CovertToString(value interface{}) string {
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
