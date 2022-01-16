package main

import (
	"fmt"

	"github.com/ankitkr1924/bankManagement/bank"
)

func main() {
	var b bank.Bank = bank.NewBank()
loop:
	for {
		fmt.Println("welcome to the Bank: ")
		fmt.Println("1. Create Your Bank Account")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Deposite Money")
		fmt.Println("4. History of transaction")
		fmt.Println("5. Delete Account")

		var choice int
		fmt.Println("Enter Your Choice : ")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			var name string
			fmt.Print("Enter Your Name: ")
			fmt.Scan(&name)
			id := b.CreateAccount(name)
			fmt.Println(id)
		case 2:
			var id string
			var amount float64
			fmt.Print("Enter Your Id  :")
			fmt.Scan(&id)
			fmt.Print("Enter 	Amount  :")
			fmt.Scanf("%f", &amount)
			b.WithDrawMoney(id, amount)
		case 3:
			var id string
			var amount float64
			fmt.Print("Enter Your Id  :")
			fmt.Scan(&id)
			fmt.Print("Enter 	Amount  :")
			fmt.Scanf("%f", &amount)
			b.DepositeMoney(id, amount)
		case 4:
			var id string
			fmt.Print("Enter Your Id  :")
			fmt.Scan(&id)
			b.GetAccountHistory(id)
		case 5:
			var id string
			fmt.Print("Enter Your Id : ")
			fmt.Scan(&id)
			b.DeleteAccount(id)
		case 6:
			fmt.Println("Thankyou")
			break loop
		default:
			fmt.Println("Invalid Choice Exiting")
			break loop
		}
	}
}
