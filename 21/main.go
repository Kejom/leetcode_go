package main

import (
	"fmt"
	"leetcode/helpers"
)

func main() {
	testCase1()
	testCase2()
	testCase3()
	testCase4()
}

func mergeTwoLists(list1 *helpers.ListNode, list2 *helpers.ListNode) *helpers.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	left := list1
	right := list2
	var mergedList = &helpers.ListNode{}
	var tail = mergedList

	for left != nil || right != nil {

		if left != nil && (right == nil || left.Val <= right.Val) {
			mergedList.Next = left
			left = left.Next
			mergedList = mergedList.Next
		} else if right != nil {
			mergedList.Next = right
			right = right.Next
			mergedList = mergedList.Next
		}

	}

	return tail.Next
}

func testCase1() {
	list1 := helpers.ListNode{Val: 1, Next: &helpers.ListNode{Val: 2, Next: &helpers.ListNode{Val: 4}}}
	list2 := helpers.ListNode{Val: 1, Next: &helpers.ListNode{Val: 3, Next: &helpers.ListNode{Val: 4}}}

	fmt.Println("Test 1")
	helpers.PrintLinkedList(mergeTwoLists(&list1, &list2))
}

func testCase2() {
	var list1 *helpers.ListNode
	var list2 *helpers.ListNode

	fmt.Println("Test 2")
	helpers.PrintLinkedList(mergeTwoLists(list1, list2))
}

func testCase3() {
	var list1 *helpers.ListNode
	list2 := helpers.ListNode{Val: 0}

	fmt.Println("Test 3")
	helpers.PrintLinkedList(mergeTwoLists(list1, &list2))
}

func testCase4() {
	list1 := helpers.ListNode{Val: 1}
	var list2 *helpers.ListNode
	fmt.Println("Test 4")
	helpers.PrintLinkedList(mergeTwoLists(&list1, list2))
}
