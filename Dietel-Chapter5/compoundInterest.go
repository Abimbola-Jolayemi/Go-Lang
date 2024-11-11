package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	years := 10

	for rate := 5; rate <= 10; rate++ {
		interestRate := float64(rate) / 100.0

		fmt.Printf("Interest Rate: %d%%\n", rate)
		fmt.Printf("%-4s %20s\n", "Year", "Amount on Deposit")

		for year := 1; year <= years; year++ {
			amount := principal * math.Pow(1.0+interestRate, float64(year))
			fmt.Printf("%-4d %20.2f\n", year, amount)
		}
		fmt.Println()
	}
}
