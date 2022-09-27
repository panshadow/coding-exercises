package testcases

var TinvertTree = []struct {
	Root  []interface{}
	Result []interface{}
}{
	{[]interface{}{4,2,7,1,3,6,9},[]interface{}{4,7,2,9,6,3,1}},
	{[]interface{}{2,1,3},[]interface{}{2,3,1}},
}
