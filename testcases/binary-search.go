package testcases

var Tsearch = []struct {
	Nums  []int
	Target int
	Result int
}{
	{[]int{-1,0,3,5,9,12}, 9, 4},
	{[]int{-1,0,3,5,9,12}, 2, -1},
	{[]int{5},5,0},
}
