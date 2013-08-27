package main

import (
	"testing"
	"strings"
	"strconv"
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
		{Query{"foo", "", 0}, baseURL + "?q=foo" + helpers + "0", false}, // Test single query
		{Query{"foo bar", "", 1}, baseURL + "?q=foo+bar" + helpers + "1", false}, // Test spaced query
		{Query{"bar baz", "go", 2}, baseURL + "?q=bar+baz+language:go" + helpers + "2", false}, // Test spaced and language
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

func TestRepoString(t *testing.T) {
	tests := []struct {
		url string
		stars int
		lang string
		name string
	}{
		{"https://github.com/foo/bar", 10, "C", "foo/bar"},
		{"othersite.com/foo/bar", 12, "C++", "othersite.com/foo/bar"},
	}

	for _, test := range tests {
		str := repoString(test.url, test.stars, test.lang)
		fields := strings.Fields(str)
		if len(fields) < 1 {
			t.Errorf("Expected fields from (%+v)", test)
		} 
		
		if fields[0] != test.name {
			t.Errorf("Expected (%s) got (%s) for (%+v)", test.name, fields[0], test)
		}

		if fields[1] != strconv.Itoa(test.stars) {
			t.Errorf("Expected (%d) got (%s) for (%+v)", test.stars, fields[1], test)
		}

		if fields[2] != test.lang {
			t.Errorf("Expected (%s) got (%s) for (%+v)", test.lang, fields[2], test)
		}
	}
}
