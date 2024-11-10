package main

import "fmt"

func main() {
	fmt.Println("How many trips were covered?")
	var trips int
	fmt.Scanln(&trips)

	var combinedMilesPerGallon float64

	for index := 0; index < trips; index++ {
		fmt.Printf("Enter miles driven for trip %d: ", index+1)
		var milesDriven float64
		fmt.Scanln(&milesDriven)

		fmt.Printf("Enter gallons of gas used for trip %d: ", index+1)
		var gasUsed float64
		fmt.Scanln(&gasUsed)

		milesPerGallon := milesDriven / gasUsed
		fmt.Printf("Miles per gallon for trip %d: %.2f\n", index+1, milesPerGallon)

		combinedMilesPerGallon += milesPerGallon
	}

	fmt.Printf("Average miles per gallon for all trips: %.2f\n", combinedMilesPerGallon/float64(trips))
}
