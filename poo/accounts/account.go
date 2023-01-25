package accounts

type Account interface {
	Deposit(value float64) bool
	Withdraw(value float64) bool
}
