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
		{Query{"", "", 0}, "", true}, // Test empty error
		{Query{"foo", "", 0}, baseURL + "?q=foo" + helpers + "10", false}, // Test single query
		{Query{"foo bar", "", 0}, baseURL + "?q=foo+bar" + helpers + "10", false}, // Test spaced query
		{Query{"bar baz", "go", 0}, baseURL + "?q=bar+baz+language:go" + helpers + "10", false}, // Test spaced and language
		{Query{"baz qux", "objc", 5}, baseURL + "?q=baz+qux+language:objc" + helpers + "5", false}, // Test custom number
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
