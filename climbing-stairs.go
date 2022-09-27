package main

import (
	"fmt"
	"leetcode/testcases"
)

var cs_memo = map[int]int{
	1: 1,
	2: 2,
}

func climbStairs(n int) int {
	res, found := cs_memo[n]
	if !found {
		if n > 2 {
			res = climbStairs(n-1) + climbStairs(n-2)
			cs_memo[n] = res
		}
	}
	return res
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TclimbStairs {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := climbStairs(t.Input)
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
