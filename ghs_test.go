package main

import (
	"testing"
)

func TestEscapeSearch(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"foo", "foo"},
		{"foo bar", "foo+bar"},
		{"foo bar baz", "foo+bar+baz"},
	}

	for _, test := range tests {
		res := escapeSearch(test.in)
		if res != test.out {
			t.Errorf("Expected (%s) for (%s) got (%s)", test.out, test.in, res)
		}
	}
}

func TestSearchString(t *testing.T) {
	tests := []struct {
		q   Query
		out string
		err bool
	}{
		{Query{"", "", 0}, "", true},
		{Query{"foo", "", 0}, baseURL + "?q=foo" + helpers + "10", false},
		{Query{"foo bar", "", 0}, baseURL + "?q=foo+bar" + helpers + "10", false},
		{Query{"bar baz", "go", 0}, baseURL + "?q=bar+baz+language:go" + helpers + "10", false},
	}

	for _, test := range tests {
		res, err := searchString(test.q)
		if test.err != (err != nil) {
			t.Errorf("Expected error to be (%t) got text (%s) for (%+v)", test.err, err.Error(), test.q)
		}

		if res != test.out {
			t.Errorf("Expected (%s) for (%+v) got (%s)", test.out, test.q, res)
		}
	}
}
