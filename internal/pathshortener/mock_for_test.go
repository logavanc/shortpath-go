package pathshortener

import "io/fs"

type (
	mockFleSystem      map[string]*mockFileSystemNode
	mockFileSystemNode struct {
		files []fs.DirEntry
		err   error
	}
	mockDirEntry struct {
		name string
	}
)

func (mfs *mockFleSystem) ReadDir(path string) (files []fs.DirEntry, err error) {
	node := (*mfs)[path]
	if node == nil {
		// If the node isn't specified, assume no other files exist at that path.
		return []fs.DirEntry{}, nil
	}
	return node.files, node.err
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
