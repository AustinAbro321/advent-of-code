package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type cool struct {
	name string
}

func TestDistance(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected int
	}{
		{
			name:     "simple",
			list1:    []int{1},
			list2:    []int{2},
			expected: 1,
		},
		{
			name:     "two numbers",
			list1:    []int{1, 2},
			list2:    []int{1, 4},
			expected: 2,
		},
		{
			name:     "two numbers",
			list1:    []int{3, 4, 2, 1, 3, 3},
			list2:    []int{4, 3, 5, 3, 9, 3},
			expected: 11,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			distance := findDistance(tc.list1, tc.list2)
			require.Equal(t, tc.expected, distance)
		})
	}
}

func TestSimilarity(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected int
	}{
		{
			name:     "simple",
			list1:    []int{1},
			list2:    []int{1},
			expected: 1,
		},
		{
			name:     "two numbers",
			list1:    []int{1, 2},
			list2:    []int{0, 2, 3},
			expected: 2,
		},
		{
			name:     "two numbers",
			list1:    []int{3, 4, 2, 1, 3, 3},
			list2:    []int{4, 3, 5, 3, 9, 3},
			expected: 31,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			distance := findSimilarity(tc.list1, tc.list2)
			require.Equal(t, tc.expected, distance)
		})
	}
}

func TestReadInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected1 []int
		expected2 []int
	}{
		{
			name: "simple",
			input: `3   4
4   3
2   5
1   3
3   9
3   3`,
			expected1: []int{3, 4, 2, 1, 3, 3},
			expected2: []int{4, 3, 5, 3, 9, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a1, a2, err := readInput(tc.input)
			require.Equal(t, tc.expected1, a1)
			require.Equal(t, tc.expected2, a2)
			require.NoError(t, err)
		})
	}
}
