package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type ListNode = h.ListNode

func showCache(cache []*ListNode) {
	for i,p := range cache {
		if p != nil {
			fmt.Printf(" {%d}/%d ",p.Val,i)
		} else {
			fmt.Printf(" {*}/%d", i)
		}
	}
	fmt.Println()
}

func roundIndexFn(n int) func (int) int {
	return func(i int) int {
		return (n+i)%n
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode  {
	icur := 0
	cur := head
	size := n+2
	rnd := roundIndexFn(size)
	cache := make([]*ListNode, size)
	num := 0
	for cur != nil {
		cache[icur]=cur
		icur = rnd(icur+1)
		num++
		cur = cur.Next
	}
	cache[icur]=nil
	bfr := cache[rnd(icur-n-1)]
	aft := cache[rnd(icur-n+1)]
	if bfr == nil {
		return aft
	}
	bfr.Next = aft
	return head
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TremoveNthFromEnd {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		head := h.UnpackListNode(t.Head)
		res := removeNthFromEnd(head,t.N)
		resList := h.PackListNode(res)
		if h.IsOK[int](resList,t.Result) {
			fmt.Println("\033[1;32mOK\033[0m", resList)
			ok++
		} else {
			fmt.Println("\033[1;31mACTUAL\033[0m", resList)
			fmt.Println("\033[1;31mEXPECTED\033[0m", t.Result)
			fail++
		}
	}
	fmt.Printf("[\033[1;32m%2d\033[0m:\033[1;31m%2d\033[0m]\033[1;37m%2d\033[0m\n", ok, fail, ok+fail)
}
