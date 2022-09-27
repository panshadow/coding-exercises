package main

import (
	"fmt"
	"leetcode/testcases"
)

func isIsomorphic(s string, t string) bool {
	tx := make(map[byte]byte)
	rx := make(map[byte]byte)
	for i,rs := range []byte(s) {
		rd, found := tx[rs]
		fmt.Printf("%d %c (%c->%c) %c \n",i,rs,rs,rd,t[i])
		if !found {
			rrs, rfound := rx[t[i]]
			if !rfound {
				tx[rs] = t[i]
				rx[t[i]]=rs
			} else if rrs != rs {
				return false
			}
		} else if rd!=t[i] {
			return false
		}
 	}

	return true
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TisIsomorphic {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := isIsomorphic(t.S, t.T)
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
