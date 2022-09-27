package testcases

var TfindDuplicate = []struct {
	Paths  []string
	Result [][]string
}{
	{[]string{
		"root/a 1.txt(abcd) 2.txt(efgh)",
		"root/c 3.txt(abcd)",
		"root/c/d 4.txt(efgh)",
		"root 4.txt(efgh)",
	}, [][]string{{
		"root/a/2.txt",
		"root/c/d/4.txt",
		"root/4.txt",
	}, {
		"root/a/1.txt",
		"root/c/3.txt",
	}}},
	{[]string{
		"root/a 1.txt(FB) 2.txt(a)",
		"root/c 3.txt(Ea)",
		"root/c/d 4.txt(b)",
		"root 4.txt(c)",
	}, [][]string{}},
}
