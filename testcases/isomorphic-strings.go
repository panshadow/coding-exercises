package testcases

var TisIsomorphic = []struct {
	S      string
	T      string
	Result bool
}{
	{"egg", "add", true},
	{"foo", "bar", false},
	{"paper", "title", true},
	{"badc","baba", false},
}
