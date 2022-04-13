package insert_delete_getRandom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	name   string
	input  input
	expect []any
}

type input struct {
	fn   []string
	nums [][]int
}

var cases = []Case{
	{
		"TestCase 1",
		input{
			[]string{"RandomizedSet", "insert", "insert", "remove", "insert", "remove", "getRandom"},
			[][]int{{}, {0}, {1}, {0}, {2}, {1}, {}},
		},
		[]any{nil, true, true, true, true, true, 2},
	},
	{
		"TestCase 2",
		input{
			[]string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
			[][]int{{}, {1}, {2}, {2}, {}, {1}, {2}, {}},
		},
		[]any{nil, true, false, true, 2, true, false, 2},
	},
	{
		"TestCase 3",
		input{
			[]string{"RandomizedSet", "remove", "remove", "insert", "getRandom", "remove", "insert"},
			[][]int{{}, {0}, {0}, {0}, {}, {0}, {0}},
		},
		[]any{nil, false, false, true, 0, true, true},
	},
}

func TestRandomizedSet(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ast.Equal(c.expect, exeFn(c.input.fn, c.input.nums), "%s", c.name)
		})
	}
}

func exeFn(fn []string, nums [][]int) []any {
	var r RandomizedSet
	var res []any
	for i, f := range fn {
		switch f {
		case "RandomizedSet":
			r = Constructor()
			res = append(res, nil)
		case "insert":
			res = append(res, r.Insert(nums[i][0]))
		case "remove":
			res = append(res, r.Remove(nums[i][0]))
		case "getRandom":
			res = append(res, r.GetRandom())
		}
	}

	return res
}
