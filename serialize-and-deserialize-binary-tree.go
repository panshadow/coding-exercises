package main

import (
	"fmt"
	h "leetcode/helpers"
	"leetcode/testcases"
	"strconv"
	"strings"
)

type TreeNode = h.TreeNode

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}


// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	curr := []*TreeNode{ root }
	next := []*TreeNode{}
	out := []string{}
	for len(curr)>0 {
		for _,node := range curr {
			if node != nil {
				out = append(out, strconv.Itoa(node.Val))
				next = append(next, node.Left, node.Right)
			} else {
				out = append(out, "null")
			}
		}
		curr = next
		next = []*TreeNode{}
	}
	return strings.Join(out, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var root *TreeNode
	svals := strings.Split(data, ",")
	if len(svals) > 0 {
		root = new(TreeNode)
		val,err := strconv.Atoi(svals[0])
		if err != nil {
			return nil
		}
		root.Val = val
		curr := []*TreeNode{ root }
		next := []*TreeNode{}
		for i:=1;i<len(svals); {
			for _, node := range curr {
				for j:=0; j<2; j++ {
					if i<len(svals) {
						if svals[i] != "null" {
							val, err := strconv.Atoi(svals[i])
							if err == nil {
								child := new(TreeNode)
								child.Val = val
								if j%2 == 0 {
									node.Left = child
								} else {
									node.Right = child
								}
								next = append(next, child)
							}
						}
						i++
					}
				}
			}
			curr = next
			next = []*TreeNode{}
		}
	}
	return root
}

func main() {
	ok := 0
	fail := 0

	for i, t := range testcases.TtreeSerializer {
		fmt.Printf("\n\n\n\033[1;36m[Test%3d]\033[0m\n", i)
		fmt.Println("Input:\n",t)
		fmt.Println("Result:")
		root := h.UnpackBTree(t.Root)
		h.ShowBTree(root)
		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(root)
		fmt.Println(data)
		res := deser.deserialize(data)
		h.ShowBTree(res)

		if h.EqBTree(res,t.Result) {
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
