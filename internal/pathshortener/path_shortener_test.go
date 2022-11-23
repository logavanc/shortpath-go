package pathshortener

import (
	"errors"
	"io/fs"
	"testing"
)

func buildMockedPathShortener() (sp *PathShortener) {
	minimumTruncationLength := 3
	truncationIndicator := '…'
	userHomePath := "/home/user"
	mfs := &mockFleSystem{
		"/home": []fs.DirEntry{
			&mockDirEntry{name: "home"},
			&mockDirEntry{name: "etc"},
			&mockDirEntry{name: "usr"},
			&mockDirEntry{name: "var"},
		},
		"/home/user": []fs.DirEntry{
			&mockDirEntry{name: "user"},
		},
		"/home/user/Downloads": []fs.DirEntry{
			&mockDirEntry{name: "Downloads"},
			&mockDirEntry{name: "Downplay"},
			&mockDirEntry{name: "Documents"},
			&mockDirEntry{name: "Desktop"},
		},
		"/home/user/Downloads/test": []fs.DirEntry{},
	}
	sp = New(
		minimumTruncationLength,
		truncationIndicator,
		userHomePath,
		mfs.ReadDir,
	)
	return
}

// func TestShortest(t *testing.T) {
// 	sp := buildMockedShortPath()

// 	testCases := []struct {
// 		name     string
// 		s        string
// 		others   []string
// 		expected string
// 	}{
// 		{"single less than ShortestLength no trunc", "a", []string{"ab", "ac"}, "a"},
// 		{"double less than ShortestLength no trunc", "ab", []string{"ab", "ac"}, "ab"},
// 		{"triple less than ShortestLength no trunc", "abc", []string{"ab", "ac"}, "abc"},
// 		{"one more than ShortestLength no trunc", "abcd", []string{"ab", "x", "abc"}, "abcd"},
// 		{"truncated at ShortestLength", "abcdefg", []string{"z", "y"}, "abc…"},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			trunc := sp.truncate(tc.s, tc.others)
// 			if trunc != tc.expected {
// 				t.Errorf("Expected %q, got %q!", tc.expected, trunc)
// 			}
// 		})
// 	}
// }

func TestShortenPath(t *testing.T) {
	ps := buildMockedPathShortener()

	testCases := []struct {
		name     string
		cwd      string
		err      error
		expected string
	}{
		{"", "", errors.New(""), unknownWorkingDir},
		{"", "/home", nil, "/home"},
		{"", "/home/user", nil, "~"},
		{"", "/home/user/Downloads", nil, "~/Downloads"},
		{"", "/home/user/Downloads/test", nil, "~/Dow…/test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sh := ps.ShortenPath(tc.cwd, tc.err)
			if sh != tc.expected {
				t.Errorf("Expected %q, got %q!", tc.expected, sh)
			}
		})
	}
}

// func BenchmarkShorten(b *testing.B) {
// 	b.ReportAllocs()
// 	sp := buildMockedPathShortener()
// 	for i := 0; i < b.N; i++ {
// 		sp.Go()
// 	}
// }

// func BenchmarkTruncate(b *testing.B) {
// 	b.ReportAllocs()

// 	minimumTruncationLength := 3
// 	truncationIndicator := '…'
// 	userHomePath := "/home/user"
// 	cwdProvider := os.Getwd
// 	filesProvider := ioutil.ReadDir

// 	sp := New(
// 		minimumTruncationLength,
// 		truncationIndicator,
// 		userHomePath,
// 		cwdProvider,
// 		filesProvider,
// 	)

// 	for i := 0; i < b.N; i++ {
// 		sp.truncate("abc", []string{"z", "y"})
// 		sp.truncate("abcdefg", []string{"z", "y"})
// 		sp.truncate("abcdefg", []string{"abcdefgh", "y"})
// 	}
// }
