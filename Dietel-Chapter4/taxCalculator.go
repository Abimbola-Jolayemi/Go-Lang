package main

import "fmt"

func main() {

	for index := 1; index <= 3; index++ {
		fmt.Printf("Enter the name of citizen %d: ", index)
		var name string
		fmt.Scan(&name)

		fmt.Printf("Enter %s's earnings: ", name)
		var earnings float64
		fmt.Scan(&earnings)

		var tax float64
		if earnings <= 30000.0 {
			tax = earnings * 0.15
		} else {
			excess := earnings - 30000.0
			tax = (30000.0 * 0.15) + (excess * 0.20)
		}

		fmt.Printf("%s's total tax: $%.2f\n", name, tax)
	}
}
