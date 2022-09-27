package testcases

var TmergeKLists = []struct {
	Lists  [][]int
	Result []int
}{
	{[][]int{
		{ 1,4,5 },{ 1,3,4 },{ 2,6 },
	}, []int{ 1,1,2,3,4,4,5,6 }},
	{[][]int{},[]int{}},
	{[][]int{{}},[]int{}},
}
