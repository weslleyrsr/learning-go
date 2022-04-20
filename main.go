package main

import (
	"fmt"
	bankAccount "learning-go/bank"
)

func main() {
	account := bankAccount.BankAccount{
		Owner:   "Jhon Doe",
		Agency:  100,
		Account: 999999,
		Balance: 10.00,
	}

	account2 := bankAccount.BankAccount{
		Owner:   "Some One",
		Agency:  100,
		Account: 777777,
		Balance: 10.00,
	}

	fmt.Println(account.Deposite(990))

	fmt.Println(account.Withdraw(90))

	fmt.Println(account.Transfer(90, &account2))
	fmt.Println(account.Balance)
	fmt.Println(account2.Balance)
}
