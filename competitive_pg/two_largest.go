package main

import "fmt"

func main() {
	i := 0
	d := []int{20, 100, 3, 4, 5, 6, 7, 8, 9, 1, 3, 20, 9, 20}
	l, sl := 0, 0

	for _, v := range d {
		i++
		if sl < v {
			i++
			if l < v {
				sl = l
				l = v
			} else {
        sl = v
      }
		}
	}

	fmt.Printf("largest: %d\nsecond largest: %d\n", l, sl)
	fmt.Printf("comparison count: %d\n", i)
}
