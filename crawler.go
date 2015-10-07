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

// Crawler is...
type Crawler struct {
	shortestLength      int
	truncationIndicator rune
}

// New returns a new Crawler instance with options configured using the
// CrawlerOption functional options.
func New(options ...CrawlerOption) (c *Crawler, err error) {
	c = &Crawler{}

	for _, option := range options {
		if err = option(c); err != nil {
			return nil, err
		}
	}

	return
}

// Shortest ...
func (c *Crawler) Shortest(s string, others []string) (short string) {
	var matched bool
	for i := c.shortestLength; i < len(s); i++ {
		matched = false
		for _, other := range others {
			if len(other) > i && s[:i] == other[:i] {
				matched = true
				break
			}
		}
		if !matched {
			return fmt.Sprintf("%s%c", s[:i], c.truncationIndicator)
		}
	}
	return s
}

// ShortPath ...
func (c *Crawler) ShortPath(dirPath string, depth int) (shortPath string) {
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
			c.ShortPath(parentPathClean, depth),
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

	shortDir := c.Shortest(dir, surroundingNodeNames)

	return fmt.Sprintf("%s%c%s",
		c.ShortPath(parentPathClean, depth),
		os.PathSeparator,
		shortDir)
}
