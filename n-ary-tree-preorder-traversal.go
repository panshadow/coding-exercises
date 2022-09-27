package main

import (
	"fmt"
	h "leetcode/helpers"
	"leetcode/testcases"
)

type Node = h.Node

func preorder(root *Node) []int {
	if root != nil {
		out := []int{root.Val}
		for _, child := range root.Children {
			out = append(out, preorder(child)...)
		}
		return out
	}
	return []int{}
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.Tpreorder {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := preorder(h.UnpackTree(t.Root))
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
