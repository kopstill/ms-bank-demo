package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (account *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	account.Balance += amount

	return nil
}

func (account *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if account.Balance < amount {
		return errors.New("the amount to withdraw is greater than the balance")
	}

	account.Balance -= amount

	return nil
}

func (account *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", account.Number, account.Name, account.Balance)
}
