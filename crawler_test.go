package main

import "testing"

func TestShortest(t *testing.T) {
	c, err := New(ShortestLength(3), TruncationIndicator('…'))
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name     string
		s        string
		others   []string
		expected string
	}{
		{"single less than ShortestLength no trunc", "a", []string{"ab", "ac"}, "a"},
		{"double less than ShortestLength no trunc", "ab", []string{"ab", "ac"}, "ab"},
		{"triple less than ShortestLength no trunc", "abc", []string{"ab", "ac"}, "abc"},
		{"one more than ShortestLength no trunc", "abcd", []string{"ab", "x", "abc"}, "abcd"},
		{"truncated at ShortestLength", "abcdefg", []string{"z", "y"}, "abc…"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shortest := c.Shortest(tc.s, tc.others)
			if shortest != tc.expected {
				t.Errorf("Expected %q, got %q!", tc.expected, shortest)
			}
		})
	}
}

func BenchmarkApp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		main()
	}
}
