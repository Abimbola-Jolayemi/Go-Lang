package main

import "fmt"

func main() {
    var num1, num2 float64

    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)

    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)

    if num1 == num2 {
        fmt.Println("0")
    } else if num1 > num2 {
        fmt.Println("1")
    } else {
        fmt.Println("-1")
    }
}
