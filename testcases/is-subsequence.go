package testcases

var TisSubsequence = []struct {
	S      string
	T      string
	Result bool
}{
	{"abc", "ahbgdc", true},
	{"axc", "ahbgdc", false},
	{"", "foobar", true},
}
