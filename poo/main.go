package main

import (
	"first-module/accounts"
	"first-module/clients"
	"log"
)

func main() {
	client1 := clients.ClientIndividual{Name: "Gustavo", Document: "777", DateBirth: "25/12/2000"}
	account1 := accounts.CurrentAccount{Client: client1, Number: 777}

	client2 := clients.ClientIndividual{Name: "Gustavo", Document: "777", DateBirth: "25/12/2000"}
	account2 := accounts.SavingAccount{Client: client2, Number: 777}

	account1.Deposit(10)
	account2.Deposit(10)

	account1.Withdraw(2)
	account2.Withdraw(2)

	account1.Transfer(3, &account2)
	account2.Transfer(1, &account1)

	log.Print(account1.GetBalance())
	log.Print(account2.GetBalance())
}
