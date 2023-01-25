package accounts

import "first-module/clients"

type SavingAccount struct {
	Client  clients.ClientIndividual
	Number  int
	balance float64
}

func (s *SavingAccount) Withdraw(value float64) bool {
	if value > s.balance && value > 0 {
		return false
	}

	s.balance -= value

	return true
}

func (s *SavingAccount) Deposit(value float64) bool {
	if value < 0 {
		return false
	}

	s.balance += value

	return true
}

func (s *SavingAccount) Transfer(value float64, to Account) bool {
	if s.balance < value {
		return false
	}

	s.balance -= value
	to.Deposit(value)
	return true
}

func (s *SavingAccount) GetBalance() float64 {
	return s.balance
}
