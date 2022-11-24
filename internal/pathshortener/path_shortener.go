package pathshortener

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

const (
	unknownWorkingDir = "???"
)

type (
	// Function signature for the os.ReadDir function.
	dirReader func(name string) ([]fs.DirEntry, error)

	PathShortener struct {
		minimumTruncationLength int
		truncationIndicator     rune
		userHomePath            string
		dirReader               dirReader
	}
)

func New(
	minimumTruncationLength int,
	truncationIndicator rune,
	userHomePath string,
	dirReader dirReader,
) (
	ps *PathShortener,
) {
	ps = &PathShortener{
		minimumTruncationLength,
		truncationIndicator,
		userHomePath,
		dirReader,
	}
	return
}

// truncate returns a truncated version of string 's' that is just long enough
// to be unique in the context of the 'others' strings. If truncation has
// occurred, the truncated string will be appended with the truncation
// indicator. The string 's' will not be truncated to shorter than the
// minimumTruncationLength.
func (ps *PathShortener) truncate(s string, others []string) (t string) {
sCharsLoop:
	for i, c := range s {
		if i < ps.minimumTruncationLength {
			t += string(c)
			continue sCharsLoop
		}
		for _, other := range others {
			if strings.HasPrefix(other, t) {
				t += string(c)
				continue sCharsLoop
			}
		}
		break sCharsLoop
	}
	if s != t {
		t += string(ps.truncationIndicator)
	}
	return
}

// Given a directory path and a file name, return a list of all the files in
// the directory that are not the file name.
func (ps *PathShortener) getOthers(
	dir string,
	name string,
) (
	others []string,
	err error,
) {
	var entries []fs.DirEntry
	entries, err = ps.dirReader(dir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != name {
			others = append(others, entry.Name())
		}
	}
	return
}

func (ps *PathShortener) shorten(p string, depth int) (short string) {
	depth++

	switch p {
	case ps.userHomePath:
		return "~"
	case "/":
		return "/"
	case "":
		return ""
	}

	parentPath, dir := path.Split(p)
	parentPathClean := parentPath[:len(parentPath)-1]

	if depth == 1 {
		return fmt.Sprintf("%s%c%s",
			ps.shorten(parentPathClean, depth),
			os.PathSeparator,
			dir,
		)
	}

	others, err := ps.getOthers(parentPathClean, dir)
	if err == nil {
		dir = ps.truncate(dir, others)
	}

	return fmt.Sprintf("%s%c%s",
		ps.shorten(parentPathClean, depth),
		os.PathSeparator,
		dir,
	)
}

func (ps *PathShortener) ShortenPath(p string, err error) (short string) {
	if err != nil {
		return unknownWorkingDir
	}
	p = path.Clean(p)
	return ps.shorten(p, 0)
}
