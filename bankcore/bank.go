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

// type Statement struct {
// 	Name    string
// 	Address string
// 	Phone   string
// 	Number  int32
// 	Balance float64
// }

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

// func (account *Account) StatementJson() Statement {
// 	return Statement{
// 		Name:    account.Name,
// 		Address: account.Address,
// 		Phone:   account.Phone,
// 		Number:  account.Number,
// 		Balance: account.Balance,
// 	}
// }

func (account *Account) Transfer(amount float64, dest *Account) error {
	if amount <= 0 {
		return errors.New("the amount to transfer should be greater than zero")
	}

	if account.Balance < amount {
		return errors.New("the amount to transfer is greater than the balance")
	}

	account.Withdraw(amount)
	dest.Deposit(amount)

	return nil
}

// // Bank ...
// type Bank interface {
// 	Statement() string
// }

// // Statement ...
// func Statement(b Bank) string {
// 	return b.Statement()
// }
