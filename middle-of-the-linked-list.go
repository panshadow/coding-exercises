package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type ListNode = h.ListNode

func middleNode(head *ListNode) *ListNode {
	index := []*ListNode{}
	cur := head
	for cur != nil {
		index = append(index, cur)
		cur = cur.Next
	}
	return index[len(index)/2]
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TmiddleNode {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := h.PackListNode(middleNode(h.UnpackListNode(t.Input)))
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
