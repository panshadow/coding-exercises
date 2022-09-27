package testcases



var TmergeTwoLists = []struct {
	List1 []int
	List2 []int
	Result []int
}{
	{[]int{1,2,4}, []int{1,3,4}, []int{1,1,2,3,4,4}},
	{[]int{}, []int{}, []int{}},
	{[]int{}, []int{0}, []int{0}},
}
