package main

import (
	"fmt"
)

// https://leetcode.com/problems/add-binary/

/*
Given two binary strings a and b, return their sum as a binary string.

Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"

Constraints:
1 <= a.length, b.length <= 10^4
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.
*/

func addBinary(a string, b string) string {
	var result string
	carry := 0

	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry == 1 {
		var sum int = carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		carry = sum / 2
		result = fmt.Sprintf("%c", sum%2+'0') + result
	}

	return result
}

func main() {
	fmt.Println(addBinary("11", "1"))      // Output: "100"
	fmt.Println(addBinary("1010", "1011")) // Output: "10101"
}
