package main

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello", "olleH"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := reverse(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
