package main

import (
	"fmt"
	h "leetcode/helpers"
	"leetcode/testcases"
	"strings"
)

func findDuplicate(paths []string) [][]string  {
	fs := make(map[string][]string)
	for _,p := range paths {
		dirInfo := strings.Split(p, " ")
		dirName := dirInfo[0]
		for _,f := range dirInfo[1:] {
			fmt.Printf("Scan files /%s/\n",f)
			fileInfo := strings.Split(f, "(")
			fileName := dirName+"/"+fileInfo[0]
			fileContent := fileInfo[1][0:len(fileInfo[1])-1]
			files,found := fs[fileContent]
			if found {
				fs[fileContent] = append(files, fileName)
			} else {
				fs[fileContent] = []string{fileName}
			}
		}
	}

	out := [][]string{}
	for fc,files := range fs {
		fmt.Println("Content",fc)
		fmt.Println("Files", files)
		if len(files) > 1 {
			out = append(out, files)
		}
	}
	return out
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TfindDuplicate {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		res := findDuplicate(t.Paths)
		if h.IsOK[string](res,t.Result) {
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
