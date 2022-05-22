package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  string
	expect int
}

var cases = []Case{
	{
		"TestCase 1",
		"abcabcbb",
		3,
	},
	{
		"TestCase 2",
		"bbbbb",
		1,
	},
	{
		"TestCase 3",
		"pwwkew",
		3,
	},
	{
		"TestCase 4",
		" ",
		1,
	},
	{
		"TestCase 5",
		"",
		0,
	},
	{
		"TestCase 5",
		"au",
		2,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, lengthOfLongestSubstring(c.input), "%s", c.name)
		})
	}
}
