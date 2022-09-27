package testcases

import (
	h "leetcode/helpers"
)

var TtreeSerializer = []struct {
	Root  []interface{}
	Result *h.TreeNode
}{
	{[]interface{}{1,2,3,nil,nil,4,5}, h.UnpackBTree([]interface{}{1,2,3,nil,nil,4,5})},
}
