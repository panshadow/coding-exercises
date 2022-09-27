package testcases

var TdetectCycle = []struct {
	Head []int
	Pos int
	Result []int
}{
	{[]int{3,2,0,-4}, 1, []int{2}},
	{[]int{1,2}, 0, []int{1}},
	{[]int{1},-1,[]int{}},
}
