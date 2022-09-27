package testcases

var TreverseList = []struct {
	Input  []int
	Result []int
}{
	{[]int{1,2,3,4,5},[]int{5,4,3,2,1}},
	{[]int{1,2}, []int{2,1}},
	{[]int{}, []int{}},
}
