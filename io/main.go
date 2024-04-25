package main

import (
	"fmt"
	"go_snippets/io/fileio"
	"os"
	"path"
)

func main() {
	args := os.Args
	dir := "data"

	fp, err := os.Open(path.Join(dir, args[1]))
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	fileio.ReadFileAndPrintReverse(fp)
}
