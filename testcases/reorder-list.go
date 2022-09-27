package testcases

var TreorderList = []struct {
	Head  []int
	Result []int
}{
	{[]int{1,2,3,4}, []int{1,4,2,3}},
	{[]int{1,2,3,4,5}, []int{1,5,2,4,3}},
}
