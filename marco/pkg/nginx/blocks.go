package nginx

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// Block A block represent a named section of an Nginx config, such as 'http', 'server' or 'location'
//  Using this object is as simple as providing a name and any sections or options,
// which can be other Block objects or option objects.
type Block struct {
	Base
	Options Directives
}

//NewBlock construct of a block
func NewBlock(name string) *Block {
	block := &Block{
		Base: NewDefaultBase(name),
	}

	block.Options = Directives{}

	return block
}

//AddDirective add options
func (b *Block) AddDirective(d Directive) {
	d.SetParent(b)
	b.Options = append(b.Options, d)
}

//AddInterface add options
func (b *Block) AddInterface(i interface{}) {
	directives, err := Marshal(i)
	if err == nil {
		for _, d := range directives {
			if d != nil {
				d.SetParent(b)
				b.Options = append(b.Options, d)
			}
		}
	}
}

//AddKVOption add options
func (b *Block) AddKVOption(key string, value interface{}) {
	d := BuildDirective(key, value)
	b.Options = append(b.Options, d)
}

// AddDirectives addall directives
func (b *Block) AddDirectives(directives []Directive) {
	if directives != nil {
		for _, d := range directives {
			b.AddDirective(d)
		}
	}
}

//Value implements Directive
func (b *Block) Value() interface{} {
	return b
}

func (b *Block) String() string {
	for _, d := range b.Options {
		d.SetIndentLevel(b.GetIndentLevel() + 1)
	}
	builder := strings.Builder{}
	for _, d := range b.Options {
		builder.WriteString(d.String())
	}

	return fmt.Sprintf("\n%s%s {%s\n%s}", b.GetIndent(), b.name, builder.String(), b.GetIndent())
}

// EmptyBlock An unnamed block of options and/or sections.
// Empty blocks are useful for representing groups of options.
type EmptyBlock struct {
	Block
}

//NewEmptyBlock new empty block
func NewEmptyBlock() *EmptyBlock {
	return &EmptyBlock{
		Block: *NewBlock(""),
	}
}

func (b *EmptyBlock) String() string {
	for _, d := range b.Options {
		d.SetIndentLevel(b.GetIndentLevel())
	}
	ds := Directives(b.Options)
	builder := strings.Builder{}
	for _, d := range ds {
		builder.WriteString(d.String())
	}

	return builder.String()
}

// Location nginx directive
type Location struct {
	Block
}

//NewLocation new location
func NewLocation(location string) *Location {
	return &Location{
		Block: *NewBlock("location " + location),
	}
}

//CustomBlock custom block like init_by_lua....
type CustomBlock struct {
	Base
	value string
}

//NewCustomBlock create custom block
func NewCustomBlock(name, value string) *CustomBlock {
	return &CustomBlock{
		Base:  NewDefaultBase(name),
		value: value,
	}
}

//Value implements Directive
func (c *CustomBlock) Value() interface{} {
	return c.value
}

func (c *CustomBlock) String() string {
	builder := strings.Builder{}
	count := c.Indent * (c.IndentLevel + 1)
	var buffer bytes.Buffer
	for i := 0; i < count; i++ {
		buffer.WriteByte(c.IndentChar)
	}
	scanner := bufio.NewScanner(strings.NewReader(c.value))
	for scanner.Scan() {
		builder.WriteString(buffer.String())
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}
	return fmt.Sprintf("\n%s%s {\n%s\n%s}", c.GetIndent(), c.name, builder.String(), c.GetIndent())
}
