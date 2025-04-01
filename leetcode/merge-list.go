package main

// https://leetcode.com/problems/merge-two-sorted-lists/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1.Next = mergeTwoLists(list1.Next, list2)
	} else {
		head = list2
		list2.Next = mergeTwoLists(list1, list2.Next)
	}

	return head
}

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, num := range nums[1:] {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return head
}

func main() {
	list1 := createList([]int{1, 2, 7})
	list2 := createList([]int{3, 4, 6})

	mergedSortedList := mergeTwoLists(list1, list2)
	head := mergedSortedList
	for head.Next != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
