package testcases

var TpseudoPalindromicPaths = []struct {
	Root  []interface{}
	Result int
}{
	{[]interface{}{2,3,1,3,1,nil,1}, 2},
	{[]interface{}{2,1,1,1,3,nil,nil,nil,nil,nil,1}, 1},
	{[]interface{}{3}, 1},
	{[]interface{}{2,1,1,1,2,nil,nil,nil,nil,nil,1}, 2},
}
