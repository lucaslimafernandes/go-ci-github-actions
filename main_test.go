package main

import (
	"testing"

	"github.com/lucaslimafernandes/go-ci-actions/utils"
)

var testValues = []struct {
	value1   int
	value2   int
	response int
}{
	{1, 1, 1},
	{1, 2, 2},
	{1, 3, 3},
	{1, 4, 4},
	{1, 5, 5},
	{1, 6, 6},
	{1, 7, 7},
	{1, 8, 8},
	{1, 9, 9},
	{1, 10, 11},
}

func TestValueGenerator(t *testing.T) {

	for _, v := range testValues {
		res := utils.ValuesGenerator(v.value1, v.value2)
		if v.response != res {
			t.Errorf("expected: %v, got: %v\n", v.response, res)
		}
	}

}
