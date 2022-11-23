// The 'shortpath' program is a small command line utility that returns a short
// but unique string representing the current working directory.
package main

import (
	"fmt"
	"os"

	"github.com/logavanc/shortpath-go/internal/pathshortener"
)

func main() {
	// TODO: Make these flags.
	shortest := 3
	indicator := '…'

	dirReader := os.ReadDir

	ps := pathshortener.New(
		shortest,
		indicator,
		os.Getenv("HOME"),
		dirReader,
	)

	shortened := ps.ShortenPath(os.Getwd())
	fmt.Printf("%s", shortened)
}
