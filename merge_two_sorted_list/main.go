package main

import "fmt"

/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

Constraints:
The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var mergedHead *ListNode
	if list1.Val < list2.Val {
		mergedHead = list1
		list1 = list1.Next
	} else {
		mergedHead = list2
		list2 = list2.Next
	}

	current := mergedHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return mergedHead
}

func main() {
	// テストケース1: list1 = [1,2,4], list2 = [1,3,4]
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergedList1 := mergeTwoLists(list1, list2)
	printList(mergedList1) // 期待される出力: [1,1,2,3,4,4]

	// テストケース2: list1 = [], list2 = []
	mergedList2 := mergeTwoLists(nil, nil)
	printList(mergedList2) // 期待される出力: []

	// テストケース3: list1 = [], list2 = [0]
	list3 := &ListNode{0, nil}
	mergedList3 := mergeTwoLists(nil, list3)
	printList(mergedList3) // 期待される出力: [0]
}

// リストの内容を表示するヘルパー関数
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
