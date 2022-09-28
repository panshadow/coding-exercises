package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

func collaps(left []int, next int) []int {
	if len(left) > 0 {
		last := left[len(left)-1]
		if last>0 && next <0 {
			if last > -next {
				return left
			} else if last == -next {
				return left[0:len(left)-1]
			} else {
				return collaps(left[0:len(left)-1], next)
			}
		}
	}
	return append(left, next)
}

func asteroidCollision(asteroids []int) []int  {
	out := []int{}
	for _,v := range asteroids {
		out = collaps(out, v)
	}
	return out
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TasteroidCollision {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		res := asteroidCollision(t.Input)
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
