package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type TreeNode = h.TreeNode

func maxNode(root *TreeNode) int {
	max := root.Val
	if root.Left != nil {
		maxLeft := maxNode(root.Left)
		if maxLeft > max {
			max = maxLeft
		}
	}

	if root.Right != nil {
		maxRight := maxNode(root.Right)
		if maxRight > max {
			max = maxRight
		}
	}
	return max
}

func minNode(root *TreeNode) int {
	min := root.Val
	if root.Left != nil {
		minLeft := minNode(root.Left)
		if minLeft < min {
			min = minLeft
		}
	}

	if root.Right != nil {
		minRight := minNode(root.Right)
		if minRight < min {
			min = minRight
		}
	}
	return min
}

func isValidBST(root *TreeNode) bool {
	var okLeft, okRight bool

	okLeft = root.Left == nil || (maxNode(root.Left) < root.Val && isValidBST(root.Left))
	okRight = root.Right == nil || (root.Val < minNode(root.Right) && isValidBST(root.Right))

	return okLeft && okRight
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TisValidBST {
		fmt.Println("Test", i)
		fmt.Println(t)
		root := h.UnpackBTree(t.Root)
		h.ShowBTree(root)
		res := isValidBST(root)
		if res == t.Result {
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
