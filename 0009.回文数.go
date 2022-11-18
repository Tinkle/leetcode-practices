package main

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var div = 1
	for x/div >= 10 {
		div *= 10
	}

	for x > 0 {
		left := x / div
		right := x % 10

		if left != right {
			return false
		}

		x = (x - left*div) / 10
		div /= 100
	}
	return true
}

// @lc code=end

// func main() {
// 	number := 145541
// 	res := isPalindrome(number)
// 	fmt.Println(res)
// }
