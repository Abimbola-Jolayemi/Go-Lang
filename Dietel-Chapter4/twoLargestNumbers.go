package main

import "fmt"


func main() {
	var largest, secondLargest int

	for index := 1; index <= 10; index++ {
		fmt.Printf("Enter number %d: ", index)
		var number int
		fmt.Scan(&number)

		if index == 1 {
			largest = number
		} else if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}
	}

	fmt.Printf("The largest number is: %d\n", largest)
	fmt.Printf("The second largest number is: %d\n", secondLargest)
}
