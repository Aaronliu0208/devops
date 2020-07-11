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
	t.Log(xType, xValue) // "[]int [1 2 3]"

	tt := &Test{5}
	tType := reflect.TypeOf(tt)
	testType := reflect.TypeOf((*Test)(nil)).Elem()
	renderType := reflect.TypeOf(new(ITRender)).Elem()
	t.Log(reflect.TypeOf(tt))

	assert.True(t, tType.Implements(renderType), "Test not implements IRender")
	assert.Equal(t, reflect.TypeOf(tt).Elem(), testType, "tt is not a type of Test")

}

func TestBaseGetIndent(t *testing.T) {
	b := NewBase("shanyou", 2, 4, 'a', nil)
	assert.Equal(t, b.GetIndent(), "aaaaaaaa", "Base getIndent fail")
}

func TestBaseRender(t *testing.T) {
	b := NewBase("server", 1, 4, 'a', nil)
	t.Log(b.Render("server"))
	assert.Equal(t, b.Render("server"), "\naaaaserver", "Base render fail")
}

func TestHash(t *testing.T) {
	var parent = "shanyou"
	b := NewBase("server", 1, 4, 'a', parent)
	t.Log(AsSha256(b))
}

func TestKeyValueOption(t *testing.T) {
	name := "shanyou"
	value := "abc"
	k := NewKeyValueOption(name, value)
	assert.Equal(t, k.Value(), "abc")

	k = NewKeyValueOption(name, &value)
	assert.Equal(t, k.Value(), "abc")

	k = NewKeyValueOption(name, false)
	assert.Equal(t, k.Value(), "off")

	k = NewKeyValueOption(name, 5)
	assert.Equal(t, k.Value(), "5")

	k = NewKeyValueOption(name, 0.5)
	assert.Equal(t, k.Value(), "0.5")

	k = NewKeyValueOption(name, []string{"a", "b", "c"})
	assert.Equal(t, k.Value(), "a b c")

	assert.Equal(t, k.String(), "\nshanyou a b c;")
}

func TestBlockRender(t *testing.T) {
	config := NewConfig()
	option := NewKeyValueOption("name", "value")
	block := NewBlock("http")
	block.AddDirective(option)
	config.AddDirective(block)
	serverblk := NewBlock("server")
	block.AddDirective(serverblk)
	serverblk.AddKVOption("listent", "80")
	location := NewLocation("/")
	location.AddKVOption("return", "200")
	serverblk.AddDirective(location)
	fmt.Println(config)
}

type TestNginxConfig struct {
	WorkerProcesses string                 `kv:"worker_processes"`
	ErrorLog        []string               `kv:"error_log,omitempty"`
	Rlimit          int                    `kv:"worker_rlimit_nofile"`
	GzipOn          bool                   `kv:"gzip"`
	Epoll           *bool                  `kv:"epoll"`
	Servers         []TestServer           `kv:"server"`
	ExtConfig       map[string]interface{} `kvext:"omitempty"`
	Marshal         *MarshalTest
}

type TestServer struct {
	Listen int    `kv:"listen"`
	Domain string `kv:"server_name"`
}

type MarshalTest struct {
}

func (t *MarshalTest) Marshal() ([]Directive, error) {
	return []Directive{
		NewKeyValueOption("marshaler", "testpass"),
	}, nil
}

func TestMarshalDirective(t *testing.T) {
	config := NewConfig()
	globalConfig := &TestNginxConfig{
		WorkerProcesses: "auto",
		Rlimit:          204800,
		GzipOn:          true,
		Servers: []TestServer{
			TestServer{
				Listen: 80,
				Domain: "www.baidu.com",
			},
		},
		ExtConfig: make(map[string]interface{}),
	}
	globalConfig.ExtConfig["abc"] = "bcd"
	globalConfig.ExtConfig["sendfile"] = "on"
	http := NewBlock("http")
	http.AddDirectives(globalConfig)
	config.AddDirective(http)
	customBlk := NewCustomBlock("init_by_lua", "kong.init()\nngx.say(ngx.var.arg_a)")
	http.AddDirective(customBlk)
	fmt.Println(config)
}

func TestMarshaler(t *testing.T) {
	test := &MarshalTest{}

	ds, err := Marshal(test)

	if err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		fmt.Println(d)
	}
}

type TestConfig struct {
	Foo     string `kv:"foo"`
	Marshal *MarshalTest
}

func TestStructFieldMarshaler(t *testing.T) {
	testConfig := &TestConfig{
		Foo:     "test",
		Marshal: &MarshalTest{},
	}

	ds, err := Marshal(testConfig)

	if err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		fmt.Println(d)
	}
}
