package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(f())

	time.Sleep(2 * time.Second)
}

func f() string {

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			g()
		}()

		h()
	}()

	return "hello"
}

func g() {
	time.Sleep(1 * time.Second)
	fmt.Println("world")
}

func h() {
	panic("panic!!!")
}
