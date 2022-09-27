package testcases

var TfloodFill = []struct {
	Image  [][]int
	Sr     int
	Sc     int
	Color  int
	Result [][]int
}{
	{[][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2, [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}},
	{[][]int{
		{0, 0, 0},
		{0, 0, 0},
	}, 0, 0, 0, [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}},
}
