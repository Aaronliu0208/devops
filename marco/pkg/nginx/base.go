package nginx

import (
	"bytes"
	"fmt"
	"strings"
)

// Directive nginx config directive
type Directive interface {
	Name() string
	Value() interface{}
	String() string
	Parent() interface{}
	SetParent(parent interface{})
	SetIndentLevel(level int)
	GetIndentLevel() int
}

// Directives slice of directive
// usage:
// sort.Sort(Directives(sliceOfDirective))
type Directives []Directive

// KVDirect represent ngin key value directive
type KVDirect = map[string]interface{}

// Len implements sort.Interface
func (d Directives) Len() int { return len(d) }

// Less implments sort.Interface
func (d Directives) Less(i, j int) bool {
	siLower := strings.ToLower(d[i].Name())
	sjLower := strings.ToLower(d[j].Name())
	if siLower == sjLower {
		return d[i].Name() < d[j].Name()
	}
	return siLower < sjLower
}

// Swap implements sort.Interface
func (d Directives) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Base nginx config item
type Base struct {
	IndentLevel int
	IndentChar  byte
	Indent      int
	parent      interface{}
	name        string
}

//NewDefaultBase default config base
func NewDefaultBase(name string) Base {
	return Base{
		IndentLevel: 0,
		IndentChar:  ' ',
		Indent:      4,
		parent:      nil,
		name:        name,
	}
}

//NewBase create new Base Object
func NewBase(name string, level, indent int, char byte, parent interface{}) Base {
	return Base{
		IndentLevel: level,
		IndentChar:  char,
		Indent:      indent,
		parent:      parent,
		name:        name,
	}
}

//GetIndent get indent for Base object
func (b Base) GetIndent() string {
	count := b.Indent * b.IndentLevel
	var buffer bytes.Buffer
	for i := 0; i < count; i++ {
		buffer.WriteByte(b.IndentChar)
	}

	return buffer.String()
}

//Render config with intent and name
func (b Base) Render(name string) string {
	return fmt.Sprintf("\n%s%s", b.GetIndent(), name)
}

// SetIndentLevel implements Directive
func (b *Base) SetIndentLevel(level int) {
	b.IndentLevel = level
}

//GetIndentLevel implements Directive
func (b Base) GetIndentLevel() int {
	return b.IndentLevel
}

//Parent implements Directive Parent
func (b Base) Parent() interface{} {
	return b.parent
}

//SetParent implements Directive
func (b *Base) SetParent(parent interface{}) {
	b.parent = parent
}

// Name implements Directive Interface
func (b Base) Name() string {
	return b.name
}
