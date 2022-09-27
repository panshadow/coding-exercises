package helpers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Any = interface{}

func UnpackListNode(xs []int) *ListNode {
	var out, last *ListNode
	for _, x := range xs {
		if out == nil {
			out = new(ListNode)
			last = out
		} else {
			last.Next = new(ListNode)
			last = last.Next
		}
		last.Val = x
	}
	return out
}

func PackListNode(list *ListNode) []int {
	out := []int{}
	cur := list
	for cur != nil {
		out = append(out, cur.Val)
		cur = cur.Next
	}
	return out
}

func ShowListNode(list *ListNode) {
	cur := list
	if cur != nil {
		for cur != nil {
			fmt.Printf("{% 3d}->",cur.Val)
			cur = cur.Next
		}
		fmt.Println("{nil}")
	}
}

func EqSlices[T comparable](list1 []T, list2 []T) bool {
	if len(list1) == len(list2) {
		for i, x := range list1 {
			if x != list2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func EqMatrix[T comparable](m1 [][]T, m2 [][]T) bool {
	if len(m1) == len(m2) {
		for i, xs := range m1 {
			if !EqSlices(xs,m2[i]) {
				return false
			}
		}
		return true
	}
	return false
}


func IsOK[T comparable](actual Any, expected Any) bool {
	switch t := actual.(type) {
	case T:
		return actual.(T)==expected.(T)
	case []T:
		return EqSlices(actual.([]T), expected.([]T))
	case [][]T:
		return EqMatrix(actual.([][]T), expected.([][]T))
	default:
		fmt.Printf("ERROR: Expected type %T", t)
	}
	return false
}


type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func EqBTree (t1 *TreeNode, t2 *TreeNode) bool {
	if t1 != nil && t2 != nil && t1.Val==t2.Val {
		return EqBTree(t1.Left, t2.Left) && EqBTree(t1.Right, t2.Right)
	} else if t1 == nil && t2 == nil {
		return true
	}
	return false
}

func SplitSliceBy(xs []interface{}) [][]int {
	out := [][]int{}
	group := []int{}
	for i := 0; i < len(xs); i++ {
		switch xs[i].(type) {
		case int:
			group = append(group, xs[i].(int))
		case nil:
			out = append(out, group)
			group = []int{}
		}
	}
	if len(group) > 0 {
		out = append(out, group)
	}
	return out
}

func UnpackTree(xs []interface{}) *Node {
	var root *Node
	if len(xs) > 0 {
		vals := SplitSliceBy(xs)
		root = new(Node)
		root.Val = vals[0][0]
		level := []*Node{root}
		nextLevel := []*Node{}
		for index := 1; index < len(vals); {
			for _, p := range level {
				if index < len(vals) {
					for _, ch := range vals[index] {
						node := new(Node)
						node.Val = ch

						p.Children = append(p.Children, node)
						nextLevel = append(nextLevel, node)
					}
					index++
				}
			}
			level = nextLevel
			nextLevel = []*Node{}
		}

	}
	return root
}

func UnpackBTree(xs []interface{}) *TreeNode {
	var root *TreeNode
	if len(xs) > 0 {
		root = new(TreeNode)
		root.Val = xs[0].(int)
		level := []*TreeNode{root}
		nextLevel := []*TreeNode{}
		for index := 1; index < len(xs); {
			for _, p := range level {
				for j := 0; j < 2; j++ {
					if index < len(xs) {
						val, ok := xs[index].(int)
						if ok {
							node := new(TreeNode)
							node.Val = val
							if j%2 == 0 {
								p.Left = node
							} else {
								p.Right = node
							}
							nextLevel = append(nextLevel, node)
						}
						index++
					}
				}
			}
			level = nextLevel
			nextLevel = []*TreeNode{}
		}
	}
	return root
}

func FindFirstNode(root *TreeNode, target int) *TreeNode {
	if root != nil {
		if root.Val == target {
			return root
		}
		find := FindFirstNode(root.Left, target)
		if find != nil {
			return find
		}
		find = FindFirstNode(root.Right, target)
		if find != nil {
			return find
		}
	}
	return nil
}

func showNode(root *Node, indent int) {
	leftPad := fmt.Sprintf("%%%ds", indent*2)
	fmt.Printf(leftPad+"[%d]\n", "+", root.Val)
	for _, ch := range root.Children {
		if ch != nil {
			showNode(ch, indent+1)
		}
	}
}

func ShowTree(root *Node) {
	showNode(root, 1)
}

func showTreeNode(root *TreeNode, showNode string, indent int) {
	if root != nil {
		leftPad := fmt.Sprintf("%%%ds", indent*2)
		fmt.Printf(leftPad+"[%d]\n", showNode, root.Val)
		showTreeNode(root.Left, "<", indent+1)
		showTreeNode(root.Right, ">", indent+1)
	}
}

func ShowBTree(root *TreeNode) {
	showTreeNode(root, ":", 1)
}
