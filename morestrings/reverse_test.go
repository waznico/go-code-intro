package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello world", "dlrow olleH"},
		{"Github", "buhtiG"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRuns(%q) == %q but want %q", c.in, got, c.want)
		}
	}
}
