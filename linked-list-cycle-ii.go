package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type ListNode = h.ListNode

func UnpackListNode(xs []int, pos int) *ListNode {
	var out, last *ListNode
	index := make([]*ListNode,len(xs))
	for i, x := range xs {
		if out == nil {
			out = new(ListNode)
			last = out
		} else {
			last.Next = new(ListNode)
			last = last.Next
		}
		index[i]= last
		last.Val = x
	}
	if pos > -1 && pos < len(index) {
		last.Next = index[pos]
	}
	return out
}

func PackListNode(list *ListNode) []int {
	if list != nil {
		return []int{list.Val}
	} else {
		return []int{}
	}
}

func detectCycle(head *ListNode) *ListNode {
	connectors := make(map[*ListNode]*ListNode)
	cur := head
	connectors[head] = nil
	for cur != nil && cur.Next != nil {
		fmt.Println("Check node ", cur.Val, " linked to ", cur.Next.Val)
		node, found := connectors[cur.Next]
		if found {
			if node != nil {
				fmt.Println("Linked from node ",node.Val, cur.Val)
			} else {
				fmt.Println("Linked from NIL node ", cur.Val)
			}
		}
		if found && node != cur {
			return cur.Next
		}
		connectors[cur.Next] = cur
		cur = cur.Next
	}
	return nil
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TdetectCycle {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := PackListNode(detectCycle(UnpackListNode(t.Head, t.Pos)))
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
