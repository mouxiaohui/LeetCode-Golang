package two_sum

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输个答案。但是，数组中同一个元素在答案里不能重复出现。
// 按任意顺序返回答案。

// 示例 1：
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

// 示例 2：
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]

// 示例 3：
// 输入：nums = [3,3], target = 6
// 输出：[0,1]

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := m[num]; ok {
			return []int{index, i}
		}
		// 以目标数与当前遍历到的数的商作为key, 遍历到的下标作为value
		m[target-num] = i
	}

	return nil
}
