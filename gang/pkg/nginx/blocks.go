package nginx

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// Dict dictionary for key value types
type Dict = map[string]Directive

// Block A block represent a named section of an Nginx config, such as 'http', 'server' or 'location'
//  Using this object is as simple as providing a name and any sections or options,
// which can be other Block objects or option objects.
type Block struct {
	Base
	Options Dict
}

// Config An unnamed block of options and/or sections.
// Empty blocks are useful for representing groups of options.
type Config struct {
	Block
}

//NewBlock construct of a block
func NewBlock(name string) *Block {
	block := &Block{
		Base: NewDefaultBase(name),
	}

	block.Options = make(Dict)

	return block
}

//AddDirective add options
func (b *Block) AddDirective(d Directive) {
	name := d.Name()
	d.SetParent(b)
	b.Options[name] = d
}

//AddDirectives add options
func (b *Block) AddDirectives(i interface{}) {
	directives, err := Marshal(i)
	if err == nil {
		for _, d := range directives {
			if d != nil {
				name := d.Name()
				d.SetParent(b)
				b.Options[name] = d
			}
		}
	}
}

//AddKVOption add options
func (b *Block) AddKVOption(key string, value interface{}) {
	d := BuildDirective(key, value)
	b.Options[key] = d
}

//Directives get all directive for blocks
func (b *Block) Directives() []Directive {
	var directives []Directive
	for _, opt := range b.Options {
		directives = append(directives, opt)
	}

	return directives
}

// SetDirectives set all directives
func (b *Block) SetDirectives(directives []Directive) {
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
	for _, d := range b.Directives() {
		d.SetIndentLevel(b.GetIndentLevel() + 1)
	}
	ds := Directives(b.Directives())
	sort.Sort(ds)
	builder := strings.Builder{}
	for _, d := range ds {
		builder.WriteString(d.String())
	}

	return fmt.Sprintf("\n%s%s {%s\n%s}", b.GetIndent(), b.name, builder.String(), b.GetIndent())
}

//NewConfig new empty block
func NewConfig() *Config {
	return &Config{
		Block: *NewBlock(""),
	}
}

func (b *Config) String() string {
	for _, d := range b.Directives() {
		d.SetIndentLevel(b.GetIndentLevel())
	}
	ds := Directives(b.Directives())
	sort.Sort(ds)
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
