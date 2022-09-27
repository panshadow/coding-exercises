package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type TreeNode = h.TreeNode

type freqMap = map[int]int

func dive(root *TreeNode, fmap freqMap, max int) (int,int) {
	if root != nil {
		nextMax := max
		sumLeft, nextMax := dive(root.Left, fmap, nextMax)
		sumRight, nextMax := dive(root.Right, fmap, nextMax)
		sum := root.Val + sumLeft + sumRight
		cnt, found := fmap[sum]
		if found {
			cnt++
		} else{
			cnt = 1
		}
		fmap[sum] = cnt
		if cnt > nextMax {
			nextMax = cnt
		}
		return sum, nextMax
	}
	return 0, max
}

func findFrequentTreeSum(root *TreeNode) []int {
	fmap := make(freqMap)
	_, max := dive(root, fmap, 0)
	out := []int{}

	for s,c := range fmap {
		if c == max {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TfindFrequentTreeSum {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		root := h.UnpackBTree(t.Root)
		res := findFrequentTreeSum(root)
		if h.IsOK[int](res, t.Result) {
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
