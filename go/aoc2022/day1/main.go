package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("in.txt")
	s := strings.Split(string(f), "\n")

	max := int64(0)
	max2 := int64(0)
	max3 := int64(0)
	tmpm := int64(0)
	for _, ss := range s {
		if ss == "" {
			if max < tmpm {
				max3 = max2
				max2 = max
				max = tmpm
			} else if max2 < tmpm {
				max3 = max2
				max2 = tmpm
			} else if max3 < tmpm {
				max3 = tmpm
			}
			tmpm = 0
		} else {
			m, _ := strconv.ParseInt(ss, 10, 64)
			tmpm += m
		}
	}

	fmt.Println(max + max2 + max3)
}
