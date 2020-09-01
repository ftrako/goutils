package goutils

import (
	"testing"
)

func  TestReadTextFile(t *testing.T) {
	n := 0
	ReadTextFile("assets/test.txt", func(text string) {
		n++
	})
	if n <= 0 {
		t.FailNow()
	}
}