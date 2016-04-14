package clg

import (
	"reflect"
	"testing"
)

func Test_Control_IfControl(t *testing.T) {
	testCases := []struct {
		Input        []interface{}
		Expected     []interface{}
		ErrorMatcher func(err error) bool
	}{
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{5, 3}, "SubtractInt", []interface{}{5, 3}},
			Expected:     []interface{}{2},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}},
			Expected:     []interface{}{},
			ErrorMatcher: nil,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}, "foo"},
			Expected:     nil,
			ErrorMatcher: IsTooManyArguments,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt"},
			Expected:     nil,
			ErrorMatcher: IsNotEnoughArguments,
		},
		{
			Input:        []interface{}{8.1, []interface{}{3, 5}, "SubtractInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", true, "SubtractInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, []int{}, []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"GreaterThanInt", []interface{}{3, 5}, "SubtractInt", true},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
		{
			Input:        []interface{}{"SplitString", []interface{}{"ab", ""}, "SubtractInt", []interface{}{5, 3}},
			Expected:     nil,
			ErrorMatcher: IsWrongArgumentType,
		},
	}

	newConfig := DefaultConfig()
	newCLGIndex, err := NewCLGIndex(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, testCase := range testCases {
		output, err := newCLGIndex.IfControl(testCase.Input...)
		if testCase.ErrorMatcher != nil && !testCase.ErrorMatcher(err) {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
		if testCase.ErrorMatcher == nil {
			if !reflect.DeepEqual(output, testCase.Expected) {
				t.Fatal("case", i+1, "expected", testCase.Expected, "got", output)
			}
		}
	}
}
