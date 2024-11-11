package main

import "fmt"

func main() {
	var binary int
	fmt.Print("Enter a binary number: ")
	fmt.Scan(&binary)

	decimal := 0
	positionalValue := 1

	for binary > 0 {
		lastDigit := binary % 10
		decimal += lastDigit * positionalValue
		positionalValue *= 2
		binary /= 10
	}

	fmt.Printf("Decimal equivalent: %d\n", decimal)
}
