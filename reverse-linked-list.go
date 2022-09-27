package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type ListNode = h.ListNode

func reverseList(head *ListNode) *ListNode {
	cur := head
	var out *ListNode
	for cur != nil {
		node := new(ListNode)
		node.Val = cur.Val
		node.Next = out
		out = node
		cur = cur.Next
	}
	return out
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TreverseList {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := h.PackListNode(reverseList(h.UnpackListNode(t.Input)))
		if h.EqSlices(res, t.Result) {
			fmt.Println("OK", res)
			ok++
		} else {
			fmt.Println("ACTUAL", res)
			fmt.Println("EXPECTED", t.Result)
			fail++
		}
		fmt.Println(ok, "/", fail, ok+fail)
	}
}
