// +build pseudo-palindromic-paths-in-a-binary-tree

package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type TreeNode = h.TreeNode

type oddMap = map[int]bool

func updateMap(m oddMap, x int) {
	odd, found := m[x]
	if found {
		m[x] = !odd
	} else {
		m[x] = true
	}
}

func isPalindron(m oddMap) bool {
	oddCounter :=0
	for _,odd := range m {
		if odd {
			if oddCounter > 0 {
				return false
			}
			oddCounter++
		}
	}
	return true
}

func getOddMaps(root *TreeNode) []oddMap {
	out := []oddMap{}
	if root.Left == nil && root.Right == nil {
		m := map[int]bool{ root.Val: true}
		return []oddMap{ m }
	}
	if root.Left != nil {
		for _,m := range getOddMaps(root.Left) {
			updateMap(m, root.Val)
			out = append(out, m)
		}
	}
	if root.Right != nil {
		for _,m := range getOddMaps(root.Right) {
			updateMap(m, root.Val)
			out = append(out, m)
		}
	}
	return out
}

func dive(m oddMap, root *TreeNode) int {
	updateMap(m, root.Val)
	defer updateMap(m, root.Val)
	if root.Left == nil && root.Right == nil {
		if isPalindron(m) {
			return 1
		}
		return 0
	}

	var res int
	if root.Left != nil {
		res += dive(m, root.Left)
	}
	if root.Right != nil {
		res += dive(m, root.Right)
	}
	return res
}

func pseudoPalindromicPathsA(root *TreeNode) int {
	counter := 0
	for _,m := range getOddMaps(root) {
		if isPalindron(m) {
			counter++
		}
	}
	return counter
}

func pseudoPalindromicPathsB(root *TreeNode) int {
	m := make(oddMap)
	return dive(m, root)
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TpseudoPalindromicPaths {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		root := h.UnpackBTree(t.Root)
		h.ShowBTree(root)
		res := pseudoPalindromicPathsB(root)
		if res == t.Result {
			fmt.Println("\033[1;32mOK\033[0m", res)
			ok++
		} else {
			fmt.Println("\033[1;31mACTUAL\033[0m", res)
			fmt.Println("EXPECTED", t.Result)
			fail++
		}
	}
	fmt.Printf("[\033[1;32m%2d\033[0m:\033[1;31m%2d\033[0m]\033[1;37m%2d\033[0m\n", ok, fail, ok+fail)
}
