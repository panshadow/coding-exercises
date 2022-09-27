package testcases

var TvalidUtf8 = []struct {
	Data  []int
	Result bool
}{
	{[]int{197,130,1}, true},
	{[]int{235,140,4}, false},
}
