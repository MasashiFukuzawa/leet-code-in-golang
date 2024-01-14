package main

import "fmt"

// https://leetcode.com/problems/sqrtx/

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.
For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

Example 1:
Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.

Example 2:
Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

Constraints:

0 <= x <= 2^31 - 1
*/

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 0, x/2

	for left <= right {
		mid := left + (right-left)/2
		squared := mid * mid

		if squared == x {
			return mid
		} else if squared < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func main() {
	fmt.Println(mySqrt(4)) // Output: 2
	fmt.Println(mySqrt(8)) // Output: 2
}
