package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)


type point = [2]int


func findIsland(grid [][]byte, p point, vmap map[point]bool) bool {
	_, found := vmap[p]
	if !found {
		vmap[p] = true
		r := p[0]
		c := p[1]
		if grid[r][c]=='1' {
			if r>0 {
				findIsland(grid,point{r-1,c}, vmap)
			}
			if r<len(grid)-1 {
				findIsland(grid,point{r+1,c}, vmap)
			}
			if c>0 {
				findIsland(grid, point{r,c-1}, vmap)
			}
			if c<len(grid[r])-1 {
				findIsland(grid, point{r,c+1}, vmap)
			}
			return true
		}
	}
	return false
}

func numIslands(grid [][]byte) int {
	visited := make(map[point]bool)
	count := 0
	for r,row := range grid {
		for c,_ := range(row) {
			p := point{r,c}
			_, found := visited[p]
			if !found {
				if findIsland(grid, p, visited) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TnumIslands {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		res := numIslands(t.Grid)
		if h.IsOK[int](res,t.Result) {
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
