package nginx

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testMapWithTag struct {
	HTTP Options `kv:"http"`
}

type testMapWithOutTag struct {
	HTTP Options
}

func TestMapWithTag(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	test := testMapWithTag{
		HTTP: Options{
			Pair{"abc", "bcd"},
		},
	}
	ds, err := Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		emptyBlk.AddDirective(d)
	}
	str := emptyBlk.String()
	fmt.Println(str)
	assert.True(t, strings.Contains(strings.ToLower(str), "http"))
	assert.True(t, strings.Contains(strings.ToLower(str), "abc bcd;"))
}

func TestMapWithOutTag(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	test := testMapWithOutTag{
		HTTP: Options{
			Pair{"abc", "bcd"},
		},
	}
	ds, err := Marshal(test)
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range ds {
		emptyBlk.AddDirective(d)
	}

	str := emptyBlk.String()
	fmt.Println(str)
	assert.False(t, strings.Contains(strings.ToLower(str), "http"))
	assert.True(t, strings.Contains(strings.ToLower(str), "abc bcd;"))
}
