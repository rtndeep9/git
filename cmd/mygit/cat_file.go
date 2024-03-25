package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var NULL_SP = []byte("\x00")

func CatFile(sha string) []byte {
	fp := filepath.Join(".git", "objects", sha[:2], sha[2:])
	contents, err := os.ReadFile(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		os.Exit(1)
	}

	zr, err := zlib.NewReader(bytes.NewReader(contents))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating zlib reader: %s\n", err)
		os.Exit(1)
	}

	defer zr.Close()

	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, zr); err != nil {
		fmt.Fprintf(os.Stderr, "Error copying data: %s\n", err)
		os.Exit(1)
	}

	return bytes.SplitN(buffer.Bytes(), NULL_SP, 2)[1]
}
