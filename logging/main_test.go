package logging

import (
	"fmt"
	"log"
	"testing"
)

func TestLogOutput(t *testing.T) {
	log.Printf("Standard Logger:")
	log.Printf("foo bar baz")
	log.Printf("foo bar baz")

	log.Printf("Long Logger:")

	// add file name and line number to standard flags
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("foo bar baz")
	log.Printf("foo bar baz")

	log.Printf("fmt.print:")

	fmt.Printf("foo bar baz")
	fmt.Printf("foo bar baz")

}
