package main

import (
	"errors"
	"fmt"
)

type E string

const (
	EUnknown E = ""
	E1       E = "e1"
	E2       E = "e2"
)

func New(s string) (E, error) {
	switch s {
	case string(E1):
		return E1, nil
	case string(E2):
		return E2, nil
	default:
		return EUnknown, errors.New(fmt.Sprintf("not found: %s", s))
	}
}

func main() {
	m1, err := New("e1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m1)

	m2, err := New("e2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(m2))
}
