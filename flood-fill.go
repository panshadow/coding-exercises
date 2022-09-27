package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type point = [2]int

func recFill(log map[point]int, image [][]int, p point, color int, target int) {
	if _,arrived := log[p]; !arrived {
		pr := p[0]
		pc := p[1]
		fmt.Printf("Fill {%d:%d} %d -> %d\n",pr,pc,target,color)
		log[p]=color
		if image[p[0]][p[1]] != color {
			image[p[0]][p[1]] = color
		}
		if pr>0 && image[pr-1][pc]==target {
			recFill(log, image, point{pr-1,pc}, color, target)
		}
		if pr<len(image)-1 && image[pr+1][pc]==target{
			recFill(log, image, point{pr+1,pc}, color, target)
		}
		if pc>0 && image[pr][pc-1]==target{
			recFill(log, image, point{pr,pc-1}, color, target)
		}
		if pc<len(image[0])-1 && image[pr][pc+1]==target{
			recFill(log, image, point{pr,pc+1}, color, target)
		}
	} else {
		fmt.Printf("---- Done {%d:%d}\n", p[0],p[1])
	}
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	targetColor := image[sr][sc]
	fillLog := make(map[point]int)
	out := make([][]int,len(image))
	for i,r := range image {
		out[i] = make([]int,len(r))
		for j,c := range r {
			out[i][j] = c
		}
	}

	recFill(fillLog, out, point{sr,sc}, color, targetColor)
	return out
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TfloodFill {
		fmt.Println("Test", i)
		fmt.Println(t)
		res := floodFill(t.Image, t.Sr, t.Sc, t.Color)
		if h.EqMatrix(res, t.Result) {
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
