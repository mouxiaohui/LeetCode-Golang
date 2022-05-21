package generate_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  int
	expect []string
}

var cases = []Case{
	{
		"TestCase 1",
		3,
		[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
	{
		"TestCase 2",
		1,
		[]string{"()"},
	},
	{
		"TestCase 3",
		2,
		[]string{"(())", "()()"},
	},
}

func TestGenerateParenthesis(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(
				c.expect,
				generateParenthesis(c.input),
				"%s", c.name,
			)
		})
	}
}
