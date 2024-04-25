package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	denom := decimal.New(12, 0)
	numer, _ := decimal.NewFromString("90000.5")

	fmt.Printf("%v", numer.DivRound(denom, 2))
}
