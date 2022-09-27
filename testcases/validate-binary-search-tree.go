package testcases

var TisValidBST = []struct {
	Root  []interface{}
	Result bool
}{
	{[]interface{}{2,1,3}, true},
	{[]interface{}{5,1,4,nil, nil, 3,6}, false},
	{[]interface{}{5,4,6,nil,nil,3,7}, false},
}
