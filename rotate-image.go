package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

func rotate(matrix [][]int)  {
	sz := len(matrix)-1
	hfv := (sz+1)/2
	hfh := hfv+(sz+1)%2
	for r:=0; r<hfv; r++ {
		for c:=0; c<hfh; c++ {
			matrix[r][c],matrix[c][sz-r],matrix[sz-r][sz-c],matrix[sz-c][r] = matrix[sz-c][r],matrix[r][c],matrix[c][sz-r],matrix[sz-r][sz-c]
		}
	}
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.Trotate {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		res := t.Matrix
		rotate(res)
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
