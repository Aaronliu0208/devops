package nginx

// Option represet nginx option
type Option interface {
	GetKey() string
	GetValue() interface{}
}

//Options list for option
type Options = []Option
