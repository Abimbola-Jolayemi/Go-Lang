package main

import (
	"fmt"
)

func main() {
	var base int

	for {
		fmt.Print("Enter the base length of the triangle (1-10): ")
		fmt.Scan(&base)
		if base >= 1 && base <= 10 {
			break
		} else {
			fmt.Println("Invalid input! Please enter a number between 1 and 10.")
		}
	}

	for row := 1; row <= base; row++ {
		for column := 1; column <= row; column++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
