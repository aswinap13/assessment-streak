package services

import (
	"errors"
	"reflect"
	"testing"
)

func TestFindPairsWithSum(t *testing.T) {
	service := NewArrayService()

	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected [][]int
		err      error
	}{
		{
			name:     "Empty array",
			numbers:  []int{},
			target:   10,
			expected: nil,
			err:      errors.New("the input array is empty"),
		},
		{
			name:     "Valid pairs",
			numbers:  []int{1, 2, 3, 4, 5},
			target:   5,
			expected: [][]int{{0, 3}, {1, 2}},
			err:      nil,
		},
		{
			name:     "No valid pairs",
			numbers:  []int{1, 2, 3},
			target:   10,
			expected: [][]int{},
			err:      nil,
		},
		{
			name:     "Array with duplicate numbers",
			numbers:  []int{1, 1, 2, 3, 4, 4, 5},
			target:   6,
			expected: [][]int{{0, 6}, {2, 5}},
			err:      nil,
		},
		{
			name:     "Negative numbers",
			numbers:  []int{-5, -3, -1, 0, 3, 5, 7},
			target:   2,
			expected: [][]int{{0, 6}, {1, 5}, {2, 4}},
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.FindPairsWithSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected result %v, got %v", tt.expected, result)
			}
		})
	}
}
