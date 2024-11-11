package main

import ( "fmt"
"math"
)

func main() {
	var numOfValues int

	fmt.Print("Enter the number of values you want to input: ")
	fmt.Scan(&numOfValues)

	if numOfValues <= 0 {
		fmt.Println("You must enter a positive number of values.")
		return
	}

	min := math.MaxInt64
	max := math.MinInt64

	fmt.Println("Enter the values:")
	for index := 0; index < numOfValues; index++ {
		var value int
		fmt.Scan(&value)

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	sum := min + max

	fmt.Printf("Minimum value: %d\n", min)
	fmt.Printf("Maximum value: %d\n", max)
	fmt.Printf("Sum of the minimum and maximum values: %d\n", sum)
}
