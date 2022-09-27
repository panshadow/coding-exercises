package main

import (
	"fmt"
	"leetcode/testcases"
)

func isSubsequence(s string, t string) bool {
	var j int
	if s == "" {
		return true
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			if j == len(s)-1 {
				return true
			} else {
				j++
			}
		}
	}
	return false
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TisSubsequence {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := isSubsequence(t.S, t.T)
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
