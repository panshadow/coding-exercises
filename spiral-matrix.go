package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

func spiralOrder(matrix [][]int) []int {
	ords := [][]int{
		{0,1},
		{1,0},
		{0,-1},
		{-1,0},
	}
	m := len(matrix[0])
	n := len(matrix)
	borders := [][]int{
		{0, n},
		{0, m},
	}
	r := 0
	c := 0
	oi := 0
	dr := ords[oi][0]
	dc := ords[oi][1]
	out := make([]int,m*n)
	for i:=0;i<m*n;i++ {
		fmt.Println("Get ",r,":",c)
		out[i] = matrix[r][c]

		if r+dr>=borders[0][0] && r+dr<borders[0][1] && c+dc>=borders[1][0] && c+dc<borders[1][1] {
			fmt.Println("Step",dr,dc)
		} else {
			oi = (oi+1)%len(ords)
			dr = ords[oi][0]
			dc = ords[oi][1]
			fmt.Println("Rotate",oi,dr,dc)
			if dr > 0 {
				fmt.Println("Change R")
				borders[0][0]++
			} else if dr< 0 {
				fmt.Println("Change L")
				borders[0][1]--
			}
			if dc > 0 {
				fmt.Println("Change B")
				borders[1][0]++
			} else if dc< 0 {
				fmt.Println("Change T")
				borders[1][1]--
			}
		}
		r += dr
		c += dc
	}
	return out
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TspiralOrder {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		res := spiralOrder(t.Matrix)
		if h.IsOK[int](res,t.Result) {
			fmt.Println("\033[1;32mOK\033[0m", res)
			ok++
		} else {
			fmt.Println("\033[1;31mACTUAL\033[0m", res)
			fmt.Println("\033[1;31mEXPECTED\033[0m", t.Result)
			fail++
		}
	}
	fmt.Printf("[\033[1;32m%2d\033[0m:\033[1;31m%2d\033[0m]\033[1;37m%2d\033[0m\n", ok, fail, ok+fail)
}
