package main

import "fmt"

func main() {
	sum := 0

	for index := 1; index <= 30; index++ {
		if index % 3 == 0 {
			sum += index
		}
	}

	fmt.Printf("The sum of integers between 1 and 30 that are divisible by 3 is: %d\n", sum)
}
