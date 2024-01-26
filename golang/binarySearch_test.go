package main


import "testing"

type BinarySearchTest struct {
	arg1     *[]int32
	arg2     int32
	expected bool
}


var binarySearchTests = []BinarySearchTest{
	{&foo, 69, true},
	{&foo, 1336, false},
	{&foo, 69420, true},
	{&foo, 69421, false},
	{&foo, 1, true},
	{&foo, 0, false},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range binarySearchTests {
		if output := binarySearch(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
