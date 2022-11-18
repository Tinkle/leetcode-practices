package main

import "fmt"

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode.cn/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (75.42%)
 * Likes:    870
 * Dislikes: 0
 * Total Accepted:    362.7K
 * Total Submissions: 480.8K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
 *
 * 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
 *
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: numRows = 5
 * 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 *
 *
 * 示例 2:
 *
 *
 * 输入: numRows = 1
 * 输出: [[1]]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1
 *
 *
 *
 *
 */

// @lc code=start
func generate(numRows int) [][]int {
	nums := make([][]int, numRows)
	nums[0] = []int{1}
	for i := 1; i < numRows; i++ {
		newNums := make([]int, i+1)
		lastNums := nums[i-1]
		var l, r = 0, i - 1
		for l-r < 2 {
			var parentLeftNo, parentRightNo = 0, 0
			if l-1 >= 0 {
				parentLeftNo = lastNums[l-1]
			}
			if r+1 < i {
				parentRightNo = lastNums[r+1]
			}
			newNums[l] = lastNums[l] + parentLeftNo
			newNums[r+1] = lastNums[r] + parentRightNo
			l++
			r--
		}
		nums[i] = newNums
	}
	return nums
}

// @lc code=end
func main() {
	nums := generate(6)
	fmt.Println(nums, "nums")
}
