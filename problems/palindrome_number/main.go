package main

import (
	"fmt"
)

/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

Constraints:
-2^31 <= x <= 2^31 - 1
*/
func isPalindrome(x int) bool {
	if x < 0 {
		// 負の数は回文になり得ない
		return false
	}

	original := x
	reversed := 0

	// 数値を逆順にする
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 逆順にした数値が元の数値と同じかチェック
	return original == reversed
}

func main() {
	result1 := isPalindrome(121)
	fmt.Println(result1) // Expected Output: true
	result2 := isPalindrome(-121)
	fmt.Println(result2) // Expected Output: false
	result3 := isPalindrome(10)
	fmt.Println(result3) // Expected Output: false
}
