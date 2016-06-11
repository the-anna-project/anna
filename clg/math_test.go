package clg

import (
	"testing"
)

func Test_CLG_Difference(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: float64(-9),
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: float64(21),
		},
		{
			A:        -3.5,
			B:        -7.5,
			Expected: float64(4),
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: float64(8),
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: float64(30),
		},
		{
			A:        11.11,
			B:        10.10,
			Expected: float64(1.0099999999999998),
		},
	}

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).Difference(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}

func Test_CLG_Sum(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: 16,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 50,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: 4,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 17,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 43,
		},
		{
			A:        36,
			B:        6.5,
			Expected: 42.5,
		},
		{
			A:        99.99,
			B:        12.15,
			Expected: 112.14,
		},
		{
			A:        17,
			B:        65,
			Expected: 82,
		},
	}

	for i, testCase := range testCases {
		output := testMaybeNewCollection(t).Sum(testCase.A, testCase.B)

		if output != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
		}
	}
}
