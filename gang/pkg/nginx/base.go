package nginx

import (
	"bytes"
	"fmt"
)

// Base nginx config base render
type Base struct {
	IndentLevel int
	IndentChar  byte
	Indent      int
	Parent      interface{}
}

//NewDefaultBase default config base
func NewDefaultBase() Base {
	return Base{
		IndentLevel: 0,
		IndentChar:  ' ',
		Indent:      4,
		Parent:      nil,
	}
}

//NewBase create new Base Object
func NewBase(level, indent int, char byte, parent interface{}) Base {
	return Base{
		IndentLevel: level,
		IndentChar:  char,
		Indent:      indent,
		Parent:      parent,
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

//Render render object
func (b Base) Render(name string) string {
	return fmt.Sprintf("\n%s%s", b.GetIndent(), name)
}
