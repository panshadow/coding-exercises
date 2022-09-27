package main

import (
	"fmt"
	"leetcode/testcases"
)

func longestPalindrome(s string) int {
	rmap := make(map[rune]int)
	for _, r := range s {
		_, found := rmap[r]
		if !found {
			rmap[r] = 1
		} else {
			rmap[r]++
		}
	}
	result := 0
	for _, count := range rmap {
		if result % 2 == 0 {
			result += count
		} else {
			result += count - count%2
		}
	}
	return result
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TlongestPalindrome {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := longestPalindrome(t.Input)
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
