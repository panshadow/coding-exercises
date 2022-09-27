package main

import (
	"fmt"
	"leetcode/testcases"
)

func runningSum(nums []int) []int {
	out := make([]int,len(nums), len(nums))
	var last int
	for i,x := range nums {
		last += x
		out[i] = last
	}
	return out
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.RunningSum {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := runningSum(t.Input)
		if Equal(res, t.Result) {
			fmt.Println("OK", res)
			ok++
		} else {
			fmt.Println("ACTUAL", res)
			fmt.Println("EXPECTED",t.Result)
			fail++
		}
		fmt.Println(ok,"/",fail,ok+fail)
	}
}
