// This file is part of the 'shortpath' program, a small command line utility
// that returns a short but unique string representing the current working
// directory.
// Copyright (C) 2015  Logan VanCuren
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software Foundation,
// Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301  USA
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	shortestLength      = 3
	truncationIndicator = '…'
)

func shortest(node string, otherNodes []string) (short string) {
	var matched bool
	for i := shortestLength; i < len(node); i++ {
		matched = false
		for _, otherNode := range otherNodes {
			if len(otherNode) > i && node[:i] == otherNode[:i] {
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Sprintf("%s%c", node[:i], truncationIndicator)
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
