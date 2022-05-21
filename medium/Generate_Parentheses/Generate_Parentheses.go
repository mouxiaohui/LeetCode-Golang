package generate_parentheses

/*
第 22 题

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

示例 3：
输入：n = 2
输出：["(())","()()"]


提示：
1 <= n <= 8
*/

func generateParenthesis(n int) []string {
	res := []string{}

	var dfs func(left int, right int, combine string)
	dfs = func(left int, right int, combine string) {
		if 2*n == len(combine) {
			res = append(res, combine)
			return
		}
		if left > 0 {
			dfs(left-1, right, combine+"(")
		}
		if left < right {
			dfs(left, right-1, combine+")")
		}
	}

	dfs(n, n, "")

	return res
}
