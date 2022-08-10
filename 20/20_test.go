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
		name:           "simple string",
		str:            "snow dog sun",
		expectedResult: "sun dog snow",
	},
	{
		name:           "many spaces",
		str:            "    snow    dog sun  ",
		expectedResult: "sun dog snow",
	},
	{
		name:           "special characters",
		str:            "sn\row  do\ng s\tun",
		expectedResult: "un s g do ow sn",
	},
	{
		name:           "single word",
		str:            "word",
		expectedResult: "word",
	},
	{
		name:           "nothing",
		str:            "",
		expectedResult: "",
	},
	{
		name:           "only special char",
		str:            "\r\n",
		expectedResult: "",
	},
}

func TestReverseWords(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r1 := ReverseWords(tc.str)
			assert.Equal(t, tc.expectedResult, r1)
		})
	}
}
