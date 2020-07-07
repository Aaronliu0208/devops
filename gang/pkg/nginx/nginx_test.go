package nginx

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ITRender interface {
	print()
}
type Test struct {
	foo int
}

func (t *Test) print() {
	fmt.Printf("%d/n", t.foo)
}

func TestRefect(t *testing.T) {
	var x interface{} = []int{1, 2, 3}
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType, xValue) // "[]int [1 2 3]"

	tt := &Test{5}
	tType := reflect.TypeOf(tt)
	testType := reflect.TypeOf((*Test)(nil)).Elem()
	renderType := reflect.TypeOf(new(ITRender)).Elem()
	fmt.Println(reflect.TypeOf(tt))

	assert.True(t, tType.Implements(renderType), "Test not implements IRender")
	assert.Equal(t, reflect.TypeOf(tt).Elem(), testType, "tt is not a type of Test")

}

func TestBaseGetIndent(t *testing.T) {
	b := NewBase(2, 4, 'a', nil)
	fmt.Println(b.GetIndent())
	assert.Equal(t, b.GetIndent(), "aaaaaaaa", "Base getIndent fail")
}

func TestBaseRender(t *testing.T) {
	b := NewBase(1, 4, 'a', nil)
	fmt.Println(b.Render("server"))
	assert.Equal(t, b.Render("server"), "\naaaaserver", "Base render fail")
}

func TestHash(t *testing.T) {
	var parent = "shanyou"
	b := NewBase(1, 4, 'a', parent)
	fmt.Println(AsSha256(b))
}

func TestAttrDicAdd(t *testing.T) {
	var parent = "shanyou"
	d := NewAttrDict(parent)
	b := NewBase(1, 4, 'a', nil)
	d.Append(&b)
	assert.Equal(t, b.Parent, parent)
}

func TestKeyOption(t *testing.T) {
	k := NewKeyOption("keyoption")
	var parent = "shanyou"
	d := NewAttrDict(parent)
	d.Append(&k)
	assert.Equal(t, k.Parent, parent)
	fmt.Println(k.String())
}

func TestKeyValueOption(t *testing.T) {
	name := "shanyou"
	value := "abc"
	k := NewKeyValueOption(name, value)
	assert.Equal(t, k.Value, "abc")

	k = NewKeyValueOption(name, &value)
	assert.Equal(t, k.Value, "abc")

	k = NewKeyValueOption(name, false)
	assert.Equal(t, k.Value, "off")

	k = NewKeyValueOption(name, 5)
	assert.Equal(t, k.Value, "5")

	k = NewKeyValueOption(name, 0.5)
	assert.Equal(t, k.Value, "0.5")

	k = NewKeyValueOption(name, []string{"a", "b", "c"})
	assert.Equal(t, k.Value, "a b c")

	assert.Equal(t, k.String(), "\nshanyou a b c;")
}

func TestKeyValuesMultiLines(t *testing.T) {
	name := "shanyou"
	values := []string{"abc", "bcd", "cde"}
	k := NewKeyValuesMultiLines(name, values)
	assert.Equal(t, k.String(), "\nshanyou abc;\nshanyou bcd;\nshanyou cde;")
}
