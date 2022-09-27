package main

import (
	"fmt"
	"leetcode/testcases"
)

var badVersion int

func isBadVersion(version int) bool {
	return version >= badVersion
}

func firstBadVersion(n int) int {
	good := 0
	bad := n
	for bad - good > 1 {
		next := (bad + good) / 2
		if isBadVersion(next) {
			fmt.Println("BAD",good,next,bad)
			bad = next
		} else {
			fmt.Println("GOOD",good,next,bad)
			good = next
		}
	}
	return bad
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TfirstBadVersion {
		fmt.Println("Test", i)
		fmt.Println(t)
		badVersion = t.Bad
		res := firstBadVersion(t.N)
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
