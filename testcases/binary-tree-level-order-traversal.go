package testcases

var TlevelOrder = []struct {
	Root  []interface{}
	Result [][]int
}{
	{[]interface{}{3,9,20,nil,nil,15,7}, [][]int{{3}, {9,20}, {15,7}} },
	{[]interface{}{1},[][]int{{1}}},
	{[]interface{}{},[][]int{}},
}
