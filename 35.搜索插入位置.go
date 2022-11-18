package main

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode.cn/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (45.02%)
 * Likes:    1781
 * Dislikes: 0
 * Total Accepted:    989.3K
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 请必须使用时间复杂度为 O(log n) 的算法。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,3,5,6], target = 5
 * 输出: 2
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,3,5,6], target = 2
 * 输出: 1
 *
 *
 * 示例 3:
 *
 *
 * 输入: nums = [1,3,5,6], target = 7
 * 输出: 4
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 为 无重复元素 的 升序 排列数组
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	var (
		start int = 0
		end   int = len(nums) - 1
		mid   int = (start + end/2)
	)

	if target < nums[start] {
		return 0
	}
	if target > nums[end] {
		return end + 1
	}

	for end >= start {
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
		mid = (start + end) / 2
	}
	return mid + 1
}

// @lc code=end

// func main() {
// 	var nums = []int{1}
// 	res := searchInsert(nums, 1)
// 	fmt.Println(res, "res")
// }
