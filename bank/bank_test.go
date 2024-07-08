package bank

import "testing"

func TestAccount(t *testing.T) {

	account := Account{
		Customer: Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("Can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("Balance is not being updated after a deposit")
	}
}

func TestDepositNegative(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("Deposit of a negative amount must return an error")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1001,
		Balance: 10,
	}

	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("Balance is not being updated after a withdrawal")
	}
}

func TestWithdrawNegative(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1001,
		Balance: 10,
	}

	if err := account.Withdraw(-10); err == nil {
		t.Error("Withdrawal of a negative amount must return an error")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()

	if statement != "1001, Jacob - Balance: 100" {
		t.Error("Statement is not being generated correctly")
	}
}

func TestAccountTransfer(t *testing.T) {
	account1 := Account{
		Customer: Customer{
			Name:    "Jacob",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1001,
		Balance: 1000,
	}

	account2 := Account{
		Customer: Customer{
			Name:    "Jane",
			Address: "1234 Main St",
			Phone:   "123-456-7890",
		},
		Number:  1002,
		Balance: 0,
	}

	account1.Transfer(&account2, 50)

	if account1.Balance != 950 || account2.Balance != 50 {
		t.Error("Transfer is not working")
	}
}
