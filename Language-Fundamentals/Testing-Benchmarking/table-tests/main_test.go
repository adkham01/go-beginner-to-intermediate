package main

import (
	"fmt"
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{data: []int{2, 3}, answer: 5},
		{data: []int{5, 9}, answer: 14},
		{data: []int{1, 2, 3, 4, 5}, answer: 15},
		{data: []int{}, answer: 0},
	}

	for _, tc := range tests {
		result := mySum(tc.data...)
		if result != tc.answer {
			t.Errorf("For input %v, expected %d but got %d", tc.data, tc.answer, result)
		}
		fmt.Println("Test passed for input:", tc.data)
	}
}
