package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type bpoint struct {
	r int
	c int
}

func hasbeen(points []bpoint, point bpoint) bool {
	for _,p := range points {
		if p.r==point.r && p.c==point.c {
			return true
		}
	}
	return false
}

func scanWord(visited []bpoint, board [][]byte, p bpoint, word string, ch int) bool {
	if !hasbeen(visited, p) {
		fmt.Printf("Check for %s[%c]%s at {%d %d} -> %c\n",word[0:ch],word[ch],word[ch+1:len(word)], p.r, p.c, board[p.r][p.c])
		if board[p.r][p.c] == byte(word[ch]) {
			fmt.Println("Found")
			if ch==len(word)-1 {
				return true
			} else {
				if p.r>0 {
					if scanWord(append(visited, p), board, bpoint{p.r-1,p.c}, word, ch+1) {
						return true
					}
				}
				if p.r<len(board)-1 {
					if scanWord(append(visited, p), board, bpoint{p.r+1,p.c}, word, ch+1) {
						return true
					}
				}
				if p.c > 0 {
					if scanWord(append(visited, p), board, bpoint{p.r,p.c-1}, word, ch+1) {
						return true
					}
				}
				if p.c < len(board[p.r])-1 {
					if scanWord(append(visited, p), board, bpoint{p.r,p.c+1}, word, ch+1) {
						return true
					}
				}
			}
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	for r,row := range board {
		for c,cell := range row {
			if cell == byte(word[0]) {
				if scanWord([]bpoint{}, board, bpoint{r,c}, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func ShowBoard(board [][]byte) {
	for r,row := range board {
		fmt.Printf("%03d |",r)
		for _, cell := range row {
			fmt.Printf("% 2c",cell)
		}
		fmt.Println("| ")
	}
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.Texist {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		ShowBoard(t.Board)
		fmt.Println("Result:")
		res := exist(t.Board, t.Word)
		if h.IsOK[bool](res,t.Result) {
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
