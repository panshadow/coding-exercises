package main

import (
	"fmt"
	"leetcode/testcases"
)

func search(nums []int, target int) int {

	from := 0
	to := len(nums)-1
	for i:=(to+from)/2; from<=to; i=(to+from)/2 {
		fmt.Printf("i=%d [%d:%d]\n", i,from,to)
		switch {
		case nums[i] == target:
			return i
		case target < nums[i]:
			to = i-1
		case target > nums[i]:
			from = i+1
		}
	}
	return -1
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.Tsearch {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := search(t.Nums, t.Target)
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
