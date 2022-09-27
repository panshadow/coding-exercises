package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type TreeNode = h.TreeNode

func EqSliceOfSlice(list1 [][]int, list2 [][]int) bool {
	if len(list1) == len(list2) {
		for i, xs := range list1 {
			if ! h.EqSlices(xs, list2[i]) {
				return false
			}
		}
		return true
	}
	return false

}

func levelOrder(root *TreeNode) [][]int {
	h.ShowBTree(root)
	out := [][]int{}
	if root != nil {
		level := []*TreeNode{root}
		nextLevel := []*TreeNode{}
		group := []int{}
		for len(level)>0 {
			for _,node := range level {
				group = append(group, node.Val)
				if node.Left != nil {
					nextLevel = append(nextLevel, node.Left)
				}
				if node.Right != nil {
					nextLevel = append(nextLevel, node.Right)
				}
			}
			out = append(out, group)
			level = nextLevel
			nextLevel = []*TreeNode{}
			group = []int{}
		}
	}

	return out
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TlevelOrder {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := levelOrder(h.UnpackBTree(t.Root))
		if EqSliceOfSlice(res, t.Result) {
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
