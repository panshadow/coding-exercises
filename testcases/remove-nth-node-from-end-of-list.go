package testcases

var TremoveNthFromEnd = []struct {
	Head  []int
	N int
	Result []int
}{
	{[]int{1,2,3,4,5}, 2, []int{1,2,3,5}},
	{[]int{1}, 1, []int{}},
	{[]int{1,2},1,[]int{1}},
	{[]int{1,2},2,[]int{2}},
}
