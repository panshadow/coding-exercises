package main

import (
	"fmt"
	"leetcode/testcases"
)

func matchFirstOctetFor(data int, n int) bool {
	mask := 0
	for i:=0;i<n;i++{
		mask = (mask<<1)|1
	}
	mask <<=1
	return data >> (7-n) ^ mask == 0
}

func matchNextOctets(data []int) bool {
	for _,b := range data {
		if b>>6 ^ 0b10 != 0 {
			return false
		}
	}
	return true
}

func validUtf8(data []int) bool {
	for i:=0; i<len(data); {
		fmt.Printf("Check %d byte: %08b\n",i,data[i])
		switch {
		case data[i]>>7 == 0:
			fmt.Printf("Match for 1: %08b\n", data[i])
			i++
		case matchFirstOctetFor(data[i], 4):
			fmt.Printf("Match for 4: %08b\n", data[i])
			next := data[i+1:i+4]
			if len(next)!=3 || !matchNextOctets(next) {
				return false
			}
			i+=4
		case matchFirstOctetFor(data[i], 3):
			fmt.Printf("Match for 3: %08b\n", data[i])
			next := data[i+1:i+3]
			if len(next)!=2 || !matchNextOctets(next) {
				return false
			}
			i+=3
		case matchFirstOctetFor(data[i], 2):
			fmt.Printf("Match for 2: %08b\n", data[i])
			next := data[i+1:i+2]
			if len(next)!=1 || !matchNextOctets(next) {
				return false
			}
			i+=2
		default:
			return false
		}
	}
	return true
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TvalidUtf8 {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := validUtf8(t.Data)
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
