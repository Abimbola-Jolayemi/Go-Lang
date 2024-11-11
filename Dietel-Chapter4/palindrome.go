package main

import "fmt"

func main() {
	var number int

	for {
		fmt.Print("Enter a five-digit integer: ")
		fmt.Scan(&number)

		if number < 10000 || number > 99999 {
			fmt.Println("Error: Please enter a five-digit integer.")
			continue
		}

		originalNumber := number

		firstDigit := number / 10000
		secondDigit := (number / 1000) % 10
		fourthDigit := (number / 10) % 10
		lastDigit := number % 10

		if firstDigit == lastDigit && secondDigit == fourthDigit {
			fmt.Printf("%d is a palindrome.\n", originalNumber)
		} else {
			fmt.Printf("%d is not a palindrome.\n", originalNumber)
		}
		break
	}
}
