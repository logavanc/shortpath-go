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

// The 'shortpath' program is a small command line utility that returns a short
// but unique string representing the current working directory.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/logavanc/shortpath-go/internal/shortpath"
)

func main() {
	myCrawler, err := shortpath.New(
		shortpath.ShortestLength(3),
		shortpath.TruncationIndicator('…'),
	)
	if err == nil {
		cwd, err := os.Getwd()
		if err == nil {
			fmt.Printf("%s", myCrawler.ShortPath(cwd, 0))
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
