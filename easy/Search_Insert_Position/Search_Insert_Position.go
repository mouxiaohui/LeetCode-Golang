package search_insert_position

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。

示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2

示例 2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

提示:
	1 <= nums.length <= 104
	-104 <= nums[i] <= 104
	nums 为 无重复元素 的 升序 排列数组
	-104 <= target <= 104
*/

// 方法一: 速度慢
/*
func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}

	return len(nums)
}
*/

// 方法二: 二分查找
func searchInsert(nums []int, target int) int {
	lenght := len(nums)
	left, right := 0, lenght-1
	res := lenght
	for left <= right {
		middleIndex := left + (right-left)/2
		if nums[middleIndex] >= target {
			res = middleIndex
			right = middleIndex - 1
		} else {
			left = middleIndex + 1
		}
	}

	return res
}
