package nginx

import (
	"fmt"
	"reflect"
)

var (
	stringType    = reflect.TypeOf("")
	stringPtrType = reflect.TypeOf((*string)(nil))
)

// Dict dictionary for key value types
type Dict = map[string]interface{}

//AttrDict dictionary for nginx config directive
type AttrDict struct {
	Owner interface{}
	Attrs Dict
}

//NewEmptyAttrDict new empty dictionary
func NewEmptyAttrDict() *AttrDict {
	return &AttrDict{
		Owner: nil,
		Attrs: make(Dict),
	}
}

//NewAttrDict new empty dictionary
func NewAttrDict(owner interface{}) AttrDict {
	return AttrDict{
		Owner: owner,
		Attrs: make(Dict),
	}
}

//Add add key value
func (a AttrDict) Add(key string, val interface{}) {
	var v reflect.Value
	if reflect.ValueOf(val).Kind() == reflect.Ptr {
		v = reflect.ValueOf(val).Elem()
	} else {
		v = reflect.ValueOf(val)
	}
	f := v.FieldByName("Parent")
	// make sure that this field is defined, and can be changed.
	if f.IsValid() && f.CanSet() {
		// set parent
		f.Set(reflect.ValueOf(a.Owner))
	}

	a.Attrs[key] = val
}

// Append append interface with name fields
func (a AttrDict) Append(item interface{}) {
	var v reflect.Value
	if reflect.ValueOf(item).Kind() == reflect.Ptr {
		v = reflect.ValueOf(item).Elem()
	} else {
		v = reflect.ValueOf(item)
	}
	f := v.FieldByName("Name")
	if f.IsValid() && f.Kind() == reflect.String && f.String() != "" {
		// set parent
		a.Add(f.String(), item)
	} else {
		a.Add(AsSha256(item), item)
	}
}

//KeyOption epresents a directive with no value.
type KeyOption struct {
	Base
	Name string
}

//NewKeyOption new keyoption
func NewKeyOption(name string) KeyOption {
	return KeyOption{
		Base: NewDefaultBase(),
		Name: name,
	}
}

func (k KeyOption) String() string {
	return k.Base.Render(fmt.Sprintf("%s;", k.Name))
}

//KeyValueOption A key/value directive. This covers most directives available for Nginx
type KeyValueOption struct {
	Base
	Name  string
	Value interface{}
}

//NewKeyValueOption init keyvalue option with key and give value
func NewKeyValueOption(name string, value interface{}) KeyValueOption {
	return KeyValueOption{
		Base:  NewDefaultBase(),
		Name:  name,
		Value: CovertToString(value),
	}
}

func (k KeyValueOption) String() string {
	return k.Base.Render(fmt.Sprintf("%s %s;", k.Name, k.Value))
}

//KeyValuesMultiLines multi lines
type KeyValuesMultiLines struct {
	Base
	Name  string
	Lines []string
}

//NewKeyValuesMultiLines create
func NewKeyValuesMultiLines(name string, value interface{}) KeyValuesMultiLines {
	var lines []string
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice {
		len := v.Len()
		for i := 0; i < len; i++ {
			val := v.Index(i).Interface()
			valString := CovertToString(val)
			lines = append(lines, valString)
		}
	} else {
		valString := CovertToString(value)
		lines = append(lines, valString)
	}

	return KeyValuesMultiLines{
		Base:  NewDefaultBase(),
		Name:  name,
		Lines: lines,
	}
}

func (k KeyValuesMultiLines) String() string {
	result := ""
	for _, v := range k.Lines {
		result += k.Base.Render(fmt.Sprintf("%s %s;", k.Name, v))
	}

	return result
}
