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

type Bank interface {
	Statement() string
}

func Statement(b Bank) string {
	return b.Statement()
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit must be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}

	if amount <= 0 {
		return errors.New("the amount to withdraw must be greater than zero")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v, %v - Balance: %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(toAccount *Account, amount float64) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}

	if err := toAccount.Deposit(amount); err != nil {
		return err
	}

	return nil
}
