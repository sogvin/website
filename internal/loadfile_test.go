package internal

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestLoadFile(t *testing.T) {
	got := LoadFile("loadfile_test.go", 3, 4)
	exp := "import (\n\t\"testing\"\n"
	assert := asserter.New(t)
	assert().Equals(got, exp)

	defer catchPanic(t)
	LoadFile("ljlkjlk", 1, 1)
}