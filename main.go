package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func shortest(node string, otherNodes []string) (short string) {
	var matched bool
	for i := 3; i < len(node); i++ {
		matched = false
		for _, otherNode := range otherNodes {
			if len(otherNode) > i && node[:i] == otherNode[:i] {
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Sprintf("%s…", node[:i])
		}
	}
	return node
}

func getShortPath(dirPath string, depth int) (shortPath string) {
	if dirPath == os.Getenv("HOME") {
		return "~"
	} else if dirPath == "/" {
		return "/"
	} else if dirPath == "" {
		return ""
	}

	depth++

	parentPath, dir := path.Split(dirPath)
	parentPathClean := parentPath[:len(parentPath)-1]

	if depth == 1 {
		return fmt.Sprintf("%s%c%s",
			getShortPath(parentPathClean, depth),
			os.PathSeparator,
			dir)
	}

	var surroundingNodeNames []string
	nodes, _ := ioutil.ReadDir(parentPath)
	for _, node := range nodes {
		if node.IsDir() && node.Name() != dir {
			surroundingNodeNames = append(surroundingNodeNames, node.Name())
		}
	}

	shortDir := shortest(dir, surroundingNodeNames)

	return fmt.Sprintf("%s%c%s",
		getShortPath(parentPathClean, depth),
		os.PathSeparator,
		shortDir)
}

func main() {
	// db, err := bolt.Open("my.db", 0600, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	if cwd, err := os.Getwd(); err == nil {
		fmt.Printf("%s", getShortPath(cwd, 0))
	} else {
		fmt.Printf("%s", err)
	}
}
