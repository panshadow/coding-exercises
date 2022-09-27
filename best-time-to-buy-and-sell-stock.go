package main

import (
	"fmt"
	"leetcode/testcases"
)

func maxProfit(price []int) int {
	pairs := [][]int{ []int{-1, -1} }
	cur := 0
	for i:=0; i<len(price);i++ {
		if pairs[cur][0] == -1 || price[i] < price[pairs[cur][0]] {
			if pairs[cur][1] != -1 && i > pairs[cur][1] {
				pairs = append(pairs,[]int{-1,-1})
				cur++
			}
			pairs[cur][0] = i
		} else if (pairs[cur][1] == -1 || price[i] > price[pairs[cur][1]]) {
			pairs[cur][1] = i
		}
	}
	profit := 0
	for _,p := range pairs {
		if p[0] >-1 && p[1] > -1 && price[p[1]] - price[p[0]] > profit {
			profit = price[p[1]] - price[p[0]]
		}
	}
	return profit
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TmaxProfit {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := maxProfit(t.Input)
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
