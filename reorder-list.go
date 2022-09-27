package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type ListNode = h.ListNode

func reorderList(head *ListNode)  {
	linkedTo := make(map[*ListNode]*ListNode)
	last := head
	cur := head
	next := cur.Next
	for last.Next != nil {
		linkedTo[last.Next] = last
		last = last.Next
	}
	for next != nil && next.Next != nil {
		fmt.Println("cur -> ",cur.Val)
		fmt.Println("last ->",last.Val)
		fmt.Println("next -> ",next.Val)
		cur.Next = last
		linkedTo[last].Next = nil
		last.Next = next
		last = linkedTo[last]
		cur = next
		next = cur.Next
	}
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TreorderList {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		head := h.UnpackListNode(t.Head)
		reorderList(head)
		res := h.PackListNode(head)
		if h.IsOK[int](res,t.Result) {
			fmt.Println("\033[1;32mOK\033[0m", res)
			ok++
		} else {
			fmt.Println("\033[1;31mACTUAL\033[0m", res)
			fmt.Println("\033[1;31mEXPECTED\033[0m", t.Result)
			fail++
		}
	}
	fmt.Printf("[\033[1;32m%2d\033[0m:\033[1;31m%2d\033[0m]\033[1;37m%2d\033[0m\n", ok, fail, ok+fail)
}
