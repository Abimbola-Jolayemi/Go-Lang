package main

import "fmt"

func main() {
	var targetSum int
	var currentSum int
	var number int

	fmt.Print("Enter the target sum: ")
	fmt.Scan(&targetSum)

	for currentSum < targetSum {

		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		currentSum += number

		fmt.Printf("Current sum is: %d\n", currentSum)
	}

	fmt.Println("Target reached or exceeded!")
}
