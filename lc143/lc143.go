package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	for head != nil {
		var last *ListNode = head.Next
		if last == nil || last.Next == nil {
			return
		}
		for last.Next.Next != nil {
			last = last.Next
		}
		next := head.Next
		head.Next = last.Next
		head.Next.Next = next
		last.Next = nil
		head = head.Next.Next
	}
}

func main() {
	head := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	pt := &head
	for pt != nil {
		fmt.Printf("%d ", pt.Val)
		pt = pt.Next
	}
	fmt.Print("\n")
	reorderList(&head)
	pt = &head
	for pt != nil {
		fmt.Printf("%d ", pt.Val)
		pt = pt.Next
	}
	fmt.Print("\n")

}
