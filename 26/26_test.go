package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	name           string
	data           string
	expectedResult bool
}

var testCases = []TestData{
	{
		name:           "1",
		data:           "abcd",
		expectedResult: true,
	},
	{
		name:           "2",
		data:           "abCdefAaf",
		expectedResult: false,
	},
	{
		name:           "3",
		data:           "aabcd",
		expectedResult: false,
	},
	{
		name:           "4",
		data:           "a A",
		expectedResult: false,
	},
	{
		name:           "5",
		data:           "__",
		expectedResult: false,
	},
	{
		name:           "6",
		data:           "\n\n",
		expectedResult: false,
	},
	{
		name:           "7",
		data:           "",
		expectedResult: true,
	},
}

func TestStingUniqueness(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r1 := IsUniqueString1(tc.data)
			r2 := IsUniqueString2(tc.data)
			r3 := IsUniqueString3(tc.data)
			assert.Equal(t, tc.expectedResult, r1)
			assert.Equal(t, tc.expectedResult, r2)
			assert.Equal(t, tc.expectedResult, r3)
		})
	}
}
