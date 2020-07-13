package nginx

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testMapWithTag struct {
	Http Options `kv:"http"`
}

type testMapWithOutTag struct {
	Http Options
}

func TestMapWithTag(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	test := testMapWithTag{
		Http: Options{
			Pair{"abc", "bcd"},
		},
	}
	ds, err := Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	emptyBlk.AddDirectives(ds)
	str := emptyBlk.String()
	fmt.Println(str)
	assert.True(t, strings.Contains(strings.ToLower(str), "http"))
	assert.True(t, strings.Contains(strings.ToLower(str), "abc bcd;"))
}

func TestMapWithOutTag(t *testing.T) {
	emptyBlk := NewEmptyBlock()
	test := testMapWithOutTag{
		Http: Options{
			Pair{"abc", "bcd"},
		},
	}
	ds, err := Marshal(test)
	if err != nil {
		t.Fatal(err)
	}
	emptyBlk.AddDirectives(ds)
	str := emptyBlk.String()
	fmt.Println(str)
	assert.False(t, strings.Contains(strings.ToLower(str), "http"))
	assert.True(t, strings.Contains(strings.ToLower(str), "abc bcd;"))
}
