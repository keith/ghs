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
	}{
		{Query{"", "", 5}, ""},
	}

	for _, test := range tests {
		res := searchString(test.q)
		if res != test.out {
			t.Errorf("Expected (%s) for (%+v) got (%s)", test.out, test.q, res)
		}
	}
}
