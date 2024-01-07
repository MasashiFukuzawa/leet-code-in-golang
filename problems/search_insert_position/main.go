package main

import "fmt"

// https://leetcode.com/problems/search-insert-position/description/

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

Constraints:
1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums contains distinct values sorted in ascending order.
-10^4 <= target <= 10^4
*/

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // Output: 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // Output: 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // Output: 4
}
