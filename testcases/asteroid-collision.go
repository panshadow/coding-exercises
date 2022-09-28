package testcases

var TasteroidCollision = []struct {
	Input  []int
	Result []int
}{
	{[]int{5, 10, -5}, []int{5, 10}},
	{[]int{8, -8}, []int{}},
	{[]int{10, 2, -5}, []int{10}},
}
