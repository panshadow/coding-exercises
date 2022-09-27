package testcases

var Tsolve = []struct {
	Board  [][]byte
	Result [][]byte
}{
	{[][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}, [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}},
	{[][]byte{{'X'}}, [][]byte{{'X'}}},

	{[][]byte{
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'X', 'O', 'X', 'X', 'O', 'O', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O'},
		{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
	}, [][]byte{
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O', 'O'},
		{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
	}},
	{[][]byte{
		{'O', 'X', 'O', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X'},
		{'O', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'O', 'O', 'O'},
	}, [][]byte{
		{'O', 'X', 'O', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X'},
		{'O', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'O', 'O', 'O'},
	}},
}
