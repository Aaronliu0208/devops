package nginx

import (
	"fmt"
	"reflect"
)

var (
	stringType    = reflect.TypeOf("")
	stringPtrType = reflect.TypeOf((*string)(nil))
)

//KeyValueOption A key/value directive. This covers most directives available for Nginx
type KeyValueOption struct {
	Base
	name  string
	value string
}

//NewKVOption init keyvalue option with key and give value
func NewKVOption(name string, value interface{}) *KeyValueOption {
	return &KeyValueOption{
		Base:  NewDefaultBase(),
		name:  name,
		value: CovertToString(value),
	}
}

func (k KeyValueOption) String() string {
	return k.Base.Render(fmt.Sprintf("%s %s;", k.name, k.value))
}

//Value implements Directive
func (k KeyValueOption) Value() interface{} {
	return k.value
}

//Value implements Directive
func (k KeyValueOption) Name() string {
	return k.name
}
