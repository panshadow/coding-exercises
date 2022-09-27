package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type point = [2]int

func findRegion(grid [][]byte, r int, c int, vmap map[point]bool, flip bool) {
	_, visited := vmap[point{r,c}]
	if !visited {
		vmap[point{r,c}]=true
		if grid[r][c] == 'O' {
			if flip {
				grid[r][c] = 'X'
			}
			fmt.Printf("-((O))- [%d:%d] %v\n",r,c,flip)
			for i:=0;i<4;i++ {
				j:=2*i+1
				r2 := r + j%3-1
				c2 := c + j/3-1
				if (r2>=0 && r2<len(grid) && c2>=0 && c2<len(grid[r])) {
					fmt.Printf("\t\t check [%d:%d] %v",r2,c2,flip)
					findRegion(grid, r2, c2, vmap, flip)
				}
			}
		} else {
			fmt.Printf("-(( ))- [%d:%d] %v\n",r,c,flip)

		}
	}
}

func solve(grid [][]byte) {
	vmap := make(map[point]bool)
	for c,_ := range grid[0] {
		findRegion(grid, 0, c, vmap, false)
		findRegion(grid, len(grid)-1, c, vmap, false)
	}
	for r,row := range grid {
		findRegion(grid, r, 0, vmap, false)
		findRegion(grid, r, len(row)-1, vmap, false)
	}

	for r:=1; r<len(grid)-1; r++ {
		for c:=1;c<len(grid[r])-1;c++ {
			findRegion(grid, r, c, vmap, true)
		}
	}
}

func showBoard(board [][]byte,title string) {
	fmt.Println(title)
	for _,row := range board {
		fmt.Println(string(row))
	}
	fmt.Println("/",title)
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.Tsolve {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		res := t.Board
		showBoard(res,"IN")
		solve(t.Board)
		showBoard(res,"OUT")
		showBoard(t.Result, "NEED")
		if h.IsOK[byte](res,t.Result) {
			fmt.Println("\033[1;32mOK\033[0m", res)
			ok++
		} else {
			fmt.Println("\033[1;31mACTUAL\033[0m", res)
			fmt.Println("EXPECTED", t.Result)
			fail++
		}
	}
	fmt.Printf("[\033[1;32m%2d\033[0m:\033[1;31m%2d\033[0m]\033[1;37m%2d\033[0m\n", ok, fail, ok+fail)
}
