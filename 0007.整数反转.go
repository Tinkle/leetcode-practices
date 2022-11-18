package main

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var result = 0

	for x != 0 {
		mod := x % 10
		result = result*10 + mod

		if result < -2147483648 || result > 2147483647 {
			return 0
		}

		x /= 10
	}
	return result
}

// @lc code=end

// func main() {
// 	res := reverse(-123)
// 	fmt.Println(res)
// }
