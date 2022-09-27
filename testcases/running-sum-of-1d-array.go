package testcases

var RunningSum = []struct{
	Input []int
	Result []int
}{
	{Input: []int{1,2,3,4}, Result: []int{1,3,6,10}},
	{[]int{1,1,1,1,1}, []int{1,2,3,4,5}},
	{[]int{3,1,2,10,1}, []int{3,4,6,16,17}},
}
