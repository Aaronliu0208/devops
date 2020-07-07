package nginx

// Block A block represent a named section of an Nginx config, such as 'http', 'server' or 'location'
//  Using this object is as simple as providing a name and any sections or options,
// which can be other Block objects or option objects.
type Block struct {
	Base
	Name     string
	Options  AttrDict
	Sections AttrDict
}

// Blocks list of blocks
type Blocks = []Block

// Section represet nginx config section
type Section = Block

// Sections list of blocks
type Sections = []Block

// EmptyBlock An unnamed block of options and/or sections.
// Empty blocks are useful for representing groups of options.
type EmptyBlock struct {
	Block
}

// Config represet nginx config struct
type Config = EmptyBlock

//NewBlock construct of a block
func NewBlock(name string, options Options, sections Blocks) Block {
	block := Block{
		Name: name,
		Base: NewDefaultBase(),
	}

	block.Options = NewAttrDict(&block)
	block.Sections = NewAttrDict(&block)

	return block
}
