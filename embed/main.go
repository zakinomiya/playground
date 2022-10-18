package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var s string

//go:embed files/exp1.txt
var ss []byte

//go:embed hello.txt
var f embed.FS

func main() {
	fmt.Printf(s, "john")
	fmt.Print(string(ss))

	data, _ := f.ReadFile("hello.txt")
	fmt.Print(string(data))
}
