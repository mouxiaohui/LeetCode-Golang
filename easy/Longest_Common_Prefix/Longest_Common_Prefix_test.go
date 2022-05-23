package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  []string
	expect string
}

var cases = []Case{
	{
		"TestCase 1",
		[]string{"flower", "flow", "flight"},
		"fl",
	},
	{
		"TestCase 2",
		[]string{"dog", "racecar", "car"},
		"",
	},
	{
		"TestCase 3",
		[]string{"aaaa"},
		"aaaa",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, longestCommonPrefix(c.input), "%s", c.name)
		})
	}
}
