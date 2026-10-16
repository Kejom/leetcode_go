package helpers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintLinkedList(start *ListNode) {
	next := start

	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
}
