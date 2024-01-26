package main

import "testing"

type LinearSearchTest struct {
	arg1     []int32
	arg2     int32
	expected bool
}

var foo = []int32{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

var linearSearchTests = []LinearSearchTest{
	{foo, 69, true},
	{foo, 1336, false},
	{foo, 69420, true},
	{foo, 69421, false},
	{foo, 1, true},
	{foo, 0, false},
}

func TestLinearSearch(t *testing.T) {
	for _, test := range linearSearchTests {
		if output := linearSearch(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
