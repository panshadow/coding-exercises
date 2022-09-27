package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type TreeNode = h.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root!=nil {
		root.Left,root.Right = invertTree(root.Right),invertTree(root.Left)
	}
	return root
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TinvertTree {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		root := h.UnpackBTree(t.Root)
		resRoot := h.UnpackBTree(t.Result)
		res := invertTree(root)
		if h.EqBTree(res,resRoot) {
			fmt.Println("\033[1;32mOK\033[0m")
			h.ShowBTree(res)
			ok++
		} else {
			fmt.Println("\033[1;31mACTUAL\033[0m")
			h.ShowBTree(res)
			fmt.Println("EXPECTED")
			h.ShowBTree(resRoot)
			fail++
		}
	}
	fmt.Printf("[\033[1;32m%2d\033[0m:\033[1;31m%2d\033[0m]\033[1;37m%2d\033[0m\n", ok, fail, ok+fail)
}
