package main

import "fmt"

func main() {
	var largest int

	for counter := 1; counter <= 10; counter++ {
		fmt.Printf("Enter number %d: ", counter)
		var number int
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}
	}

	fmt.Printf("The largest number is: %d\n", largest)
}
