package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type ListNode = h.ListNode


func mergeKLists(lists []*ListNode) *ListNode  {
	cs := make([]*ListNode, len(lists))
	for i,l := range lists {
		cs[i] = l
	}
	next := 0
	var last,out *ListNode
	for next != -1 {
		next = -1
		for i,c := range cs {
			if c!= nil {
				if next==-1 || cs[next].Val > c.Val {
					next = i
				}
			}
		}
		if next != -1 {
			if out == nil {
				out = new(ListNode)
				last = out
			} else {
				last.Next = new(ListNode)
				last = last.Next
			}
			last.Val = cs[next].Val
			cs[next]=cs[next].Next
		}
	}
	return out
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TmergeKLists {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		lists := make([]*ListNode,len(t.Lists))
		for i,l := range t.Lists {
			lists[i] = h.UnpackListNode(l)
		}
		resList := mergeKLists(lists)
		h.ShowListNode(resList)
		res := h.PackListNode(resList)
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
