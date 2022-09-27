package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"

)

type ListNode = h.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var out, last *ListNode
	cur1 := list1
	cur2 := list2
	for cur1 != nil || cur2 != nil {
		if out == nil {
			out = new(ListNode)
			last = out
		} else {
			last.Next = new(ListNode)
			last = last.Next
		}
		if cur1 != nil && cur2 != nil {
			if cur1.Val <= cur2.Val {
				last.Val = cur1.Val
				cur1 = cur1.Next
			} else {
				last.Val = cur2.Val
				cur2 = cur2.Next
			}
		} else if cur1 != nil {
			last.Val = cur1.Val
			cur1 = cur1.Next

		} else if cur2 != nil {
			last.Val = cur2.Val
			cur2 = cur2.Next
		}
	}
	return out
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TmergeTwoLists {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := h.PackListNode(mergeTwoLists(h.UnpackListNode(t.List1), h.UnpackListNode(t.List2)))
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
