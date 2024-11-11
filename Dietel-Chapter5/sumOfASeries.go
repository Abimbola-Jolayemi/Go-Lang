package main

import "fmt"

func main() {
	fmt.Printf("%-5s %-15s\n", "n", "Sum")

	var sum int64 = 0
	for index := 1; index <= 100; index++ {
		sum += int64(index)
		fmt.Printf("%-5d %-15d\n", index, sum)
	}
}
