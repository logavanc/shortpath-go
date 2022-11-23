package pathshortener

import "io/fs"

type mockFleSystem map[string][]fs.DirEntry

func (mfs *mockFleSystem) ReadDir(path string) (files []fs.DirEntry, err error) {
	return (*mfs)[path], nil
}
