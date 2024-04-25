package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n)))
	for i := 3; i <= root; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	until := 1000000

	n := time.Now()

	fmt.Println(2)
	for i := 3; i < until; i += 2 {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

	fmt.Println(time.Since(n))
}
