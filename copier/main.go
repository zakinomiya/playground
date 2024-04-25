package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type T struct {
	Arr []*T
}

type S struct {
	Arr  []*T
	ArrB []*T
}

func main() {
	t := T{Arr: make([]*T, 10)}
	var s S
	copier.Copy(&s, t)
	fmt.Println(t, s)
	fmt.Println(s.ArrB)
	fmt.Println(len(s.ArrB))
	for range s.ArrB {
		fmt.Println("hello")
	}
}
