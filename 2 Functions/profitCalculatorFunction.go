package main

import (
	"errors"
	"fmt"
	"os"
)

func main () {

	funcRevenue, errRev := getRevenue("Enter Revenue: ")

	if(errRev != nil){
		fmt.Println(errRev)
		return
	}


	funcExpense, errExp := getExpense("Enter Expense: ")

	if(errExp != nil){
		fmt.Println(errExp)
		return
	}

	funcTax, errTax := getTax("Enter Tax-rate: ")

	if(errTax != nil){
		fmt.Println(errTax)
		return
	}

	
	// if(errRev != nil || errExp != nil || errTax != nil){
	// 	fmt.Println(errTax)
	// 	return
	// }

	fmt.Println("Revenue: ", funcRevenue)

	fmt.Println("Expense:", funcExpense)

	fmt.Println("Tax-rate: ", funcTax)


	ebtVal, profitVal, ratioVal := calculate(funcRevenue, funcExpense, funcTax)

	fmt.Printf(`EBT Value: %0.1f
Profit Value: %0.1f
Ratio Value: %0.1f`, ebtVal, profitVal, ratioVal)


		//pass data to the function and write into the file
	writeResults(ebtVal, profitVal, ratioVal)

}

func getRevenue(revMessage string) (float64, error){
	var revenue float64

	fmt.Print(revMessage)
	fmt.Scan(&revenue)

	if(revenue <= 0){
		return 0, errors.New("revenue must be grater than 0")
	}

	return revenue, nil
}

func getExpense(expenseMessage string) (float64, error){
	var expense float64

	fmt.Print(expenseMessage)
	fmt.Scan(&expense)

	if(expense <= 0){
		return 0, errors.New("expense must be grater than 0")
	}

	return expense, nil 
}

func getTax(taxMessage string) (float64, error){
	var taxRate float64

	fmt.Print(taxMessage)
	fmt.Scan(&taxRate)

	if(taxRate <= 0){
		return 0, errors.New("taxRate must be grater than 0")
	}

	return taxRate, nil 
}

func calculate (revenue float64, expense float64, taxRate float64) (float64, float64, float64 ){
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	return ebt, profit, ratio
}

func writeResults(ebt, profit, ratio float64) {

	results := fmt.Sprintf("EBT : %.1f\nProfit : %.1f\nRatio : %.1f\n", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}