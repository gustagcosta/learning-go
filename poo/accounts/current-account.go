package accounts

import "first-module/clients"

type CurrentAccount struct {
	Client  clients.ClientIndividual
	Number  int
	balance float64
}

func (c *CurrentAccount) Withdraw(value float64) bool {
	if value > c.balance && value > 0 {
		return false
	}

	c.balance -= value

	return true
}

func (c *CurrentAccount) Deposit(value float64) bool {
	if value < 0 {
		return false
	}

	c.balance += value

	return true
}

func (c *CurrentAccount) Transfer(value float64, to Account) bool {
	if c.balance < value {
		return false
	}

	c.balance -= value
	to.Deposit(value)
	return true
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
