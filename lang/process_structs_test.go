package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

type cleanPathT struct {
	Old   string
	New   string
	Clean string
}

func TestCleanPath(t *testing.T) {
	tests := []cleanPathT{
		{
			Old:   "/foo",
			New:   "bar",
			Clean: "/foo/bar",
		},
		{
			Old:   "/foo",
			New:   ".bar",
			Clean: "/foo/.bar",
		},
		{
			Old:   "/foo",
			New:   "./bar",
			Clean: "/foo/bar",
		},
		{
			Old:   "/foo",
			New:   "../bar",
			Clean: "/bar",
		},
		{
			Old:   "/foo",
			New:   "/bar",
			Clean: "/bar",
		},
		/////
		{
			Old:   "foo",
			New:   "bar",
			Clean: "foo/bar",
		},
		{
			Old:   "foo",
			New:   ".bar",
			Clean: "foo/.bar",
		},
		{
			Old:   "foo",
			New:   "./bar",
			Clean: "foo/bar",
		},
		{
			Old:   "foo",
			New:   "../bar",
			Clean: "bar",
		},
		{
			Old:   "foo",
			New:   "/bar",
			Clean: "/bar",
		},
		///// with space suffix /////
		{
			Old:   "/foo ",
			New:   "bar",
			Clean: "/foo /bar",
		},
		{
			Old:   "/foo ",
			New:   ".bar",
			Clean: "/foo /.bar",
		},
		{
			Old:   "/foo ",
			New:   "./bar",
			Clean: "/foo /bar",
		},
		{
			Old:   "/foo ",
			New:   "../bar",
			Clean: "/bar",
		},
		{
			Old:   "/foo ",
			New:   "/bar",
			Clean: "/bar",
		},
		/////
		{
			Old:   "foo ",
			New:   "bar",
			Clean: "foo /bar",
		},
		{
			Old:   "foo ",
			New:   ".bar",
			Clean: "foo /.bar",
		},
		{
			Old:   "foo ",
			New:   "./bar",
			Clean: "foo /bar",
		},
		{
			Old:   "foo ",
			New:   "../bar",
			Clean: "bar",
		},
		{
			Old:   "foo ",
			New:   "/bar",
			Clean: "/bar",
		},
		///// with space prefix /////
		{
			Old:   "/foo",
			New:   " bar",
			Clean: "/foo/ bar",
		},
		{
			Old:   "/foo",
			New:   " .bar",
			Clean: "/foo/ .bar",
		},
		{
			Old:   "/foo",
			New:   " ./bar",
			Clean: "/foo/ ./bar",
		},
		{
			Old:   "/foo",
			New:   " ../bar",
			Clean: "/foo/ ../bar",
		},
		{
			Old:   "/foo",
			New:   " /bar",
			Clean: "/foo/ /bar",
		},
		/////
		{
			Old:   "foo",
			New:   " bar",
			Clean: "foo/ bar",
		},
		{
			Old:   "foo",
			New:   " .bar",
			Clean: "foo/ .bar",
		},
		{
			Old:   "foo",
			New:   " ./bar",
			Clean: "foo/ ./bar",
		},
		{
			Old:   "foo",
			New:   " ../bar",
			Clean: "foo/ ../bar",
		},
		{
			Old:   "foo",
			New:   " /bar",
			Clean: "foo/ /bar",
		},
	}

	count.Tests(t, len(tests))

	for i := range tests {
		clean := cleanPath(tests[i].Old, tests[i].New)
		if clean != tests[i].Clean {
			t.Error("cleanPath doesn't return expected output:")
			t.Logf("  Test:      %d", i)
			t.Logf("  Expected: '%s'", tests[i].Clean)
			t.Logf("  Actual:   '%s'", clean)
			t.Logf("  Old path: '%s'", tests[i].Old)
			t.Logf("  New path: '%s'", tests[i].New)
		}
	}
}
