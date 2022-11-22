package main

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode.cn/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (68.77%)
 * Likes:    440
 * Dislikes: 0
 * Total Accepted:    240.2K
 * Total Submissions: 349K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
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
 * 输入: rowIndex = 3
 * 输出: [1,3,3,1]
 *
 *
 * 示例 2:
 *
 *
 * 输入: rowIndex = 0
 * 输出: [1]
 *
 *
 * 示例 3:
 *
 *
 * 输入: rowIndex = 1
 * 输出: [1,1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 0
 *
 *
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(rowIndex) 空间复杂度吗？
 *
 */

// @lc code=start
func getRow(rowIndex int) []int {
	var slice = []int{1}
	if rowIndex == 0 {
		return slice
	}
	for i := 0; i < rowIndex; i++ {
		slice = append(slice, 1)
		dst := make([]int, len(slice))
		copy(dst, slice)
		for i := 1; i < len(dst)-1; i++ {
			slice[i] = dst[i-1] + dst[i]
		}
	}
	return slice
}

// @lc code=end

// func main() {
// 	fmt.Println(getRow(4))
// }
