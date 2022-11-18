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

package shortpath

import "errors"

// CrawlerOption is...
type CrawlerOption func(*Crawler) error

// ShortestLength is a CrawlerOption that sets the minimum number of runes a
// directory will be shortened to.
func ShortestLength(l int) (f func(*Crawler) error) {
	f = func(c *Crawler) (err error) {
		if l > 0 {
			c.shortestLength = l
		} else {
			err = errors.New("Invalid shortest length option, must be > 0.")
		}
		return
	}
	return
}

// TruncationIndicator is a CrawlerOption that sets the runes used to indicate
// that truncation has occurred.
func TruncationIndicator(r rune) (f func(*Crawler) error) {
	f = func(c *Crawler) (err error) {
		c.truncationIndicator = r
		return
	}
	return
}
