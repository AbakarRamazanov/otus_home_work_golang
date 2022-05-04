package main

// import (
// 	"github.com/stretchr/testify/require"
// 	"testing"
// )

// func TestReverse(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{"Hello", "olleH"},
// 		{"OTUS", "SUTO"},
// 		{"Hello, OTUS!", "!SUTO ,olleH"},
// 		{"abcd", "dcba"},
// 		{" 123 =", "= 321 "},
// 		// {"+-*/", "/*-+"},
// 	}

// 	for _, tc := range tests {
// 		tc := tc
// 		t.Run(tc.input, func(t *testing.T) {
// 			result := reverse(tc.input)
// 			require.Equal(t, tc.expected, result)
// 		})
// 	}
// }
