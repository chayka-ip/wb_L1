package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	name          string
	slice         []int
	index         uint64
	expectedSlice []int
}

var testCases = []TestData{
	{
		name:          "general usage",
		slice:         []int{0, 1, 2, 3, 4, 5},
		index:         3,
		expectedSlice: []int{0, 1, 2, 4, 5},
	},
	{
		name:          "slice with one element",
		slice:         []int{0},
		index:         0,
		expectedSlice: []int{},
	},
	{
		name:          "remove first elem",
		slice:         []int{0, 1},
		index:         0,
		expectedSlice: []int{1},
	},
	{
		name:          "remove last elem",
		slice:         []int{0, 1},
		index:         1,
		expectedSlice: []int{0},
	},
}

func TestRemoveSlicePreserveOrder(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeItemPreserveOrder(tc.slice, tc.index)
			assert.Equal(t, tc.expectedSlice, res)
		})
	}
}

func TestRemoveSliceBreakOrder(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			initialSlice := append([]int{}, tc.slice...)
			res := removeItemBreakOrder(tc.slice, tc.index)
			// initial slice should not be modified
			assert.Equal(t, initialSlice, tc.slice)
			sort.Ints(tc.expectedSlice)
			sort.Ints(res)
			assert.Equal(t, tc.expectedSlice, res)
		})
	}
}

func TestRemoveSliceBreakOrderInplace(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeItemBreakOrderInplace(tc.slice, tc.index)
			sort.Ints(tc.expectedSlice)
			sort.Ints(res)
			assert.Equal(t, tc.expectedSlice, res)
		})
	}
}
