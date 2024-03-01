package main

import "fmt"

func main () {

	//declare variables
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expense: ")
	fmt.Scan(&expense)

	fmt.Print("Enter Tax-rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	// fmt.Println("EBT : ", ebt)
	// fmt.Println("Profit : ", profit)
	// fmt.Println("Ratio : ", ratio)

	//other print options
	//fmt.Printf("EBT %v\n Profit %v\n Ration %v", ebt, profit, ratio)
	//fmt.Printf("\nEBT %0.0f\nProfit %.1f\nRation %.1f", ebt, profit, ratio)

	//another way printing values using sprint
	// reEbt := fmt.Sprintf("EBT %0.0f\n", ebt)
	// reProfit := fmt.Sprintf("Profit %.1f\n", profit)
	// reRatio := fmt.Sprintf("Ration %.1f\n", ratio)

	// fmt.Print(reEbt,reProfit,reRatio)


	// fmt.Println(reEbt)
	// fmt.Println(reProfit)
	// fmt.Println(reRatio)

	//print multiline string
	fmt.Printf(`EBT %0.0f
Profit %.1f
Ration %.1f`, ebt, profit, ratio)


}