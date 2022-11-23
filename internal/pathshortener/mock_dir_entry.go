package pathshortener

import "io/fs"

type mockDirEntry struct {
	name string
}

func (mde *mockDirEntry) Name() string {
	return mde.name
}

func (mde *mockDirEntry) IsDir() bool {
	return true
}

func (mde *mockDirEntry) Type() fs.FileMode {
	return fs.ModeDir
}

func (mde *mockDirEntry) Info() (fs.FileInfo, error) {
	return nil, nil
}
