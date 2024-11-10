package main

import "fmt"

func main() {
    var accountNumber string
    var beginningBalance float64
	var totalCharges float64
	var credits float64
	var creditLimit float64


    fmt.Print("Enter account number: ")
    fmt.Scanln(&accountNumber)

    fmt.Print("Enter beginning balance: ")
    fmt.Scanln(&beginningBalance)

    fmt.Print("Enter total charges for the month: ")
    fmt.Scanln(&totalCharges)

    fmt.Print("Enter total credits applied for the month: ")
    fmt.Scanln(&credits)

    fmt.Print("What is the allowed credit limit for this customer? ")
    fmt.Scanln(&creditLimit)

    newBalance := beginningBalance + totalCharges - credits

    fmt.Printf("The new balance for this customer is: %.2f\n", newBalance)

    if newBalance > creditLimit {
        fmt.Println("Credit limit exceeded!!!")
    }
}
