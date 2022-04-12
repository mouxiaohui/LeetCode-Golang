package number_of_lines_to_write_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  input
	expect []int
}

type input struct {
	widths []int
	s      string
}

var cases = []Case{
	{
		"TesCase 1",
		input{
			[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"abcdefghijklmnopqrstuvwxyz",
		},
		[]int{3, 60},
	},
	{
		"TesCase 2",
		input{
			[]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"bbbcccdddaaa",
		},
		[]int{2, 4},
	},
}

func TestNumberOfLines(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, numberOfLines(c.input.widths, c.input.s), "%s", c.name)
		})
	}
}
