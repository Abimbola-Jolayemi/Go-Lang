package main

import "fmt"

func main() {
    var x1, y1, x2, y2 float64

    fmt.Print("Enter x-coordinate of the first point: ")
    fmt.Scan(&x1)
    fmt.Print("Enter y-coordinate of the first point: ")
    fmt.Scan(&y1)

    fmt.Print("Enter x-coordinate of the second point: ")
    fmt.Scan(&x2)
    fmt.Print("Enter y-coordinate of the second point: ")
    fmt.Scan(&y2)

    if x1 == x2 {
        fmt.Println("The line between the points is perpendicular to the y-axis.")
    } else if y1 == y2 {
        fmt.Println("The line between the points is perpendicular to the x-axis.")
    } else {
        fmt.Println("The line between the points is not perpendicular to any axis.")
    }
}
