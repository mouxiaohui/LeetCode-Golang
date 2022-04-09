package house_robber_2

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

示例 1：
输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

示例 2：
输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

示例 3：
输入：nums = [1,2,3]
输出：3

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/

// 题解: 使用动态规划
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	_rob := func(nums []int) int {
		if len(nums) == 2 {
			return max(nums[0], nums[1])
		}

		// 前两个最优解
		opt1 := nums[0]
		opt2 := max(nums[0], nums[1])

		// 对于每个房屋可以选择偷与不偷，两种选择中利益最大的为最优解。
		// 每个最优解都由前两个最优解决定。
		for i := 2; i < len(nums); i++ {
			A := nums[i] + opt1 // 选择偷
			B := opt2           // 选择不偷
			opt1 = opt2
			opt2 = max(A, B)
		}

		return opt2
	}

	// 因为首位相连，偷了第一个房屋就无法偷最后一个，所以要分两种情况: 要么偷第一个，要么偷最后个。
	return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
