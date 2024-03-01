package main

import (
	"fmt"
	"math"
)


func main() {

	
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	//investmentAmount := 1000.0
	var investmentAmount float64;

	//var expectedReturnRate float64 = 5.5
	var expectedReturnRate float64; //just declare it 
	//expectedReturnRate := 5.5
	//var years float64 = 10
	var years float64;
	

	const inflationRate = 2.5

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ");
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Years: ");
	fmt.Scan(&years)

	//type conversions
	//var futureValues = float64(investmentAmount) * math.Pow((1 + expectedReturnRate/100), float64(years))
	// var futureValues = (investmentAmount) * math.Pow((1 + expectedReturnRate/100), years)
	futureValues := (investmentAmount) * math.Pow((1 + expectedReturnRate/100), years)
	futureRealValues := futureValues / math.Pow((1+ inflationRate/100), years)

	fmt.Println(futureValues)
	fmt.Println(futureRealValues)

}