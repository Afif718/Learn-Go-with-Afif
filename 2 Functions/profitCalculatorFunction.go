package main

import "fmt"

func main () {

	funcRevenue := getRevenue("Enter Revenue: ")
	funcExpense := getExpense("Enter Expense: ")
	funcTax := getTax("Enter Tax-rate: ")

	fmt.Println("Revenue: ", funcRevenue)

	fmt.Println("Expense:", funcExpense)

	fmt.Println("Tax-rate: ", funcTax)


	ebtVal, profitVal, ratioVal := calculate(funcRevenue, funcExpense, funcTax)

	fmt.Printf(`EBT Value: %0.1f
Profit Value: %0.1f
Ratio Value: %0.1f`, ebtVal, profitVal, ratioVal)

}

func getRevenue(revMessage string) float64{
	var revenue float64

	fmt.Print(revMessage)
	fmt.Scan(&revenue)

	return revenue 
}

func getExpense(expenseMessage string) float64{
	var expense float64

	fmt.Print(expenseMessage)
	fmt.Scan(&expense)

	return expense 
}

func getTax(taxMessage string) float64{
	var taxRate float64

	fmt.Print(taxMessage)
	fmt.Scan(&taxRate)

	return taxRate 
}

func calculate (revenue float64, expense float64, taxRate float64) (float64, float64, float64 ){
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	return ebt, profit, ratio
}