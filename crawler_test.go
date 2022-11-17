package main

import "testing"

func TestShortest(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		others   []string
		expected string
	}{
		// {"test0", "abcdefg", []string{"z", "y"}, "a…"},
		// {"test1", "abcd", []string{"ab", "x"}, "abc…"},
		{"test2", "abcd", []string{"ab", "x", "abc"}, "abcd"},
		// {"test3", "a", []string{"ab", "ac"}, "a"},
	}

	c, err := New(ShortestLength(3), TruncationIndicator('…'))
	if err != nil {
		t.Fatal(err)
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

	}
}
