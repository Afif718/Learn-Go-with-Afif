package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//make the file name global
var fileName = "balance.txt"

func readDataFromFile() (float64, error){
	data, err := os.ReadFile(fileName)

	if(err != nil){
		return 100, errors.New("failed to read file")
	}


	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if(err != nil){
		return 100, errors.New("failed to read balance")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64){

	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func main(){

	//account balance variable 
	
	accountBalance, err := readDataFromFile();
	fmt.Println("Welcome to SystemSage Bank!")

	if(err != nil){
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("------------------")
		//return
		//panic("Can't Continue!")
	}

	//for i:= 0; i< 4; i++ {} this is for loop in go

	//for {} this is infinit loop

	for  {
		
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	
		//get input from user
		var choice int
		fmt.Print("Your choice for input: ")
		fmt.Scan(&choice)
	
		//check if choice is equal to 1 it will return either true or false
		// wantsCheckBalance := choice == 1

		/*
		Alternative of for loop
		switch choice {
		case: 1
			//do something
		case: 2
			//do something

		default: 
			//for invalid cases
			return // to get out the switch case

		}
		*/
	
		if(choice == 1){
			fmt.Println("Your Account Balance: ", accountBalance)
		} else if (choice ==2 ){
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if(depositAmount <=0){
				fmt.Println("Invalid Amount! Amount should be grater than 0")
				fmt.Println()
				//return 
				continue
			}
	
			accountBalance += depositAmount
	
			fmt.Println("Balance Updated! New Balance: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if (choice ==3){
			fmt.Print("Withdraw Money: ")
			var withdrawMoney float64
			fmt.Scan(&withdrawMoney)
	
			if(withdrawMoney <=0){
				fmt.Println("Invalid Amount! Amount should be grater than 0")
				fmt.Println()
				//return 
				continue
			}
	
			if(withdrawMoney > accountBalance){
				fmt.Println("Invalid Amount! You can not withdraw more than you have!")
				//return 
				continue
			}
	
			 accountBalance -= withdrawMoney
	
			fmt.Println("New Balance: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if(choice ==4) {
			fmt.Println("Goodbye!!")
			//return
			break
		} else {
			fmt.Println("Please enter from 1 to 4 only")
			fmt.Scan(&choice)
		}
	}

	fmt.Println("Thank you for choosing our bank!")

}