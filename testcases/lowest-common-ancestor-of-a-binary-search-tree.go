package testcases

var TlowestCommonAncestor = []struct {
	Root  []interface{}
	P int
	Q int
	Result int
}{
	{[]interface{}{6,2,8,0,4,7,9,nil,nil,3,5}, 2, 8, 6},
	{[]interface{}{6,2,8,0,4,7,9,nil,nil,3,5}, 2, 4, 2},
	{[]interface{}{2,1}, 2, 1, 2},
}
