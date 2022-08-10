package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	name           string
	str            string
	expectedResult string
}

var testCases = []TestData{
	{
		name:           "1",
		str:            "Hello, 世界界",
		expectedResult: "界界世 ,olleH",
	},
	{
		name:           "2",
		str:            "你好-;'|?好",
		expectedResult: "好?|';-好你",
	},
}

func TestReverseString(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r1 := ReverseString(tc.str)
			r2 := ReverseString2(tc.str)
			r3 := ReverseString3(tc.str)
			r4 := ReverseString4(tc.str)
			assert.Equal(t, tc.expectedResult, r1)
			assert.Equal(t, tc.expectedResult, r2)
			assert.Equal(t, tc.expectedResult, r3)
			assert.Equal(t, tc.expectedResult, r4)
		})
	}

}
