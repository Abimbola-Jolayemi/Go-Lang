package main

import "fmt"

func main() {
	item1Value := 239.99
	item2Value := 129.75
	item3Value := 99.95
	item4Value := 350.89

	var item1Sold, item2Sold, item3Sold, item4Sold int

	fmt.Print("Enter the number of Item 1 sold (Value: $239.99): ")
	fmt.Scan(&item1Sold)

	fmt.Print("Enter the number of Item 2 sold (Value: $129.75): ")
	fmt.Scan(&item2Sold)

	fmt.Print("Enter the number of Item 3 sold (Value: $99.95): ")
	fmt.Scan(&item3Sold)

	fmt.Print("Enter the number of Item 4 sold (Value: $350.89): ")
	fmt.Scan(&item4Sold)

	totalSales := float64(item1Sold)*item1Value + float64(item2Sold)*item2Value +
		float64(item3Sold)*item3Value + float64(item4Sold)*item4Value

	earnings := 200.0 + (totalSales * 0.09)

	fmt.Printf("The salesperson's earnings for the week: $%.2f\n", earnings)
}
