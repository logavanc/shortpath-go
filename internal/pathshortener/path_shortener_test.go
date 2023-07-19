package pathshortener

import (
	"errors"
	"io/fs"
	"testing"
)

var testCases = []struct {
	name     string
	cwd      string
	err      error
	expected string
}{
	{"root special case", "/", nil, "/"},
	{"home special case", "/home/user", nil, "~"},
	{"unknown cwd", "", errors.New(""), unknownWorkingDir},
	{"", "/home/user/aaaxxx/test", nil, "~/aaax…/test"},
	{"", "/home/user/bbbxxx/test", nil, "~/bbb…/test"},
	{"", "/secret/aaaxxx/test", nil, "/sec…/aaaxxx/test"},
	{"path cleaning", "////", nil, "/"},
}

func buildMockedPathShortener() (sp *PathShortener) {
	minimumTruncationLength := 3
	truncationIndicator := '…'
	userHomePath := "/home/user"
	mfs := &mockFleSystem{
		"/home/user": &mockFileSystemNode{
			files: []fs.DirEntry{
				&mockDirEntry{name: "aaaxxx"},
				&mockDirEntry{name: "aaayyy"},
				&mockDirEntry{name: "bbbxxx"},
			}},
		"/secret": &mockFileSystemNode{
			err: fs.ErrPermission,
			files: []fs.DirEntry{
				&mockDirEntry{name: "aaaxxx"},
				&mockDirEntry{name: "aaayyy"},
			}},
	}
	sp = New(
		minimumTruncationLength,
		truncationIndicator,
		userHomePath,
		mfs.ReadDir,
	)
	return
}

func TestShortenPath(t *testing.T) {
	ps := buildMockedPathShortener()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sh := ps.ShortenPath(tc.cwd, tc.err)
			if sh != tc.expected {
				t.Errorf("Expected %q, got %q!", tc.expected, sh)
			}
		})
	}
}

func BenchmarkShortenPath(b *testing.B) {
	ps := buildMockedPathShortener()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			ps.ShortenPath(tc.cwd, tc.err)
		}
	}
}
