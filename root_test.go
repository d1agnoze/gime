package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gabriel-vasile/mimetype"
)

func TestInputHanlder(t *testing.T) {
	t.Run("Test mimetype library: Text", func(t *testing.T) {
		file, err := os.Open("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		mtype, err := mimetype.DetectReader(file)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(mtype.String())
		if !strings.Contains(mtype.String(), "text") {
			t.Errorf("Expected text/* or application/*, got: %s", mtype.String())
		}
	})

	t.Run("Test absolute file path", func(t *testing.T) {
		file, err := os.Open("input.txt")
		if err != nil {
			t.Fatal(err)
		}
		abs, err := filepath.Abs(file.Name())
    fmt.Println(abs)
    if err != nil{
      t.Error(err)
    }
	})
}
