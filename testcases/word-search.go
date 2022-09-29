package testcases

var Texist = []struct {
	Board  [][]byte
	Word   string
	Result bool
}{
	{[][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED", true},
	{[][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE", true},
	{[][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, "cdba", true},
	{[][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCESEEEFS", true},
}
