package main

import (
	"fmt"
	"leetcode/testcases"
	h "leetcode/helpers"
)

type TreeNode = h.TreeNode

func buildUpLinks(parent,root *TreeNode, tab map[*TreeNode]*TreeNode) {
	if root!=nil {
		tab[root] = parent
		if root.Left != nil {
			buildUpLinks(root, root.Left, tab)
		}
		if root.Right != nil {
			buildUpLinks(root, root.Right, tab)
		}
	}
}

func upLinkPath(uplinks map[*TreeNode]*TreeNode, node *TreeNode) []*TreeNode {
	cur := node
	path := []*TreeNode{cur}
	for cur != nil {
		prn, found := uplinks[cur]
		if found && prn != nil {
			path = append(path, prn)
		}
		cur = prn
	}
	return path
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	upLinks := make(map[*TreeNode]*TreeNode)
	buildUpLinks(nil, root, upLinks)
	pPath := upLinkPath(upLinks, p)

	qPath := upLinkPath(upLinks, q)

	if len(pPath) > len(qPath) {
		pPath = pPath[len(pPath)-len(qPath):len(pPath)]
	} else if len(pPath) < len(qPath) {
		qPath = qPath[len(qPath)-len(pPath):len(qPath)]
	}
	for i,node := range pPath {
		if node==qPath[i] {
			return node
		}
	}

	return root
}

func main() {
	ok := 0
	fail := 0
	for i, t := range testcases.TlowestCommonAncestor {
		fmt.Println("Test", i)
		fmt.Println(t)
		root := h.UnpackBTree(t.Root)
		pNode := h.FindFirstNode(root, t.P)
		qNode := h.FindFirstNode(root, t.Q)
		h.ShowBTree(root)
		res := lowestCommonAncestor(root, pNode, qNode)
		if res.Val == t.Result {
			fmt.Println("OK", res.Val)
			ok++
		} else {
			fmt.Println("ACTUAL", res.Val)
			fmt.Println("EXPECTED", t.Result)
			fail++
		}
		fmt.Println(ok, "/", fail, ok+fail)
	}
}
