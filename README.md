# Bank Package
## Overview

The bank package provides a simple banking system that includes customer management, account operations, and basic transaction functionalities. It supports operations such as deposit, withdrawal, balance inquiry, and funds transfer between accounts.

## Installation

To use the bank package, simply import it in your Go project:

``` go
import "github.com/Jacob-Krueckl/bankcore"
```

## Package Contents

### Types

#### Customer

A Customer represents a bank customer and contains the following fields:

    Name: The name of the customer.
    Address: The address of the customer.
    Phone: The phone number of the customer.

#### Account

An Account represents a bank account and embeds a Customer. It includes the following fields:

    Customer: The customer associated with the account.
    Number: The account number.
    Balance: The current balance of the account.

### Interfaces

#### Bank

The Bank interface defines a method for generating account statements:

```go
type Bank interface {
    Statement() string
}
```

### Functions

#### Statement

The Statement function generates a statement for any type that implements the Bank interface:

```go
func Statement(b Bank) string
```

### Methods

#### Account Methods

- `Deposit(amount float64) error`
    - Deposits a specified amount into the account. Returns an error if the amount is less than or equal to zero.

- `Withdraw(amount float64) error`
    - Withdraws a specified amount from the account. Returns an error if the amount is greater than the account balance or less than or equal to zero.

- `Statement() string`
    - Generates a statement string for the account.

- `Transfer(toAccount *Account, amount float64) error`
    - Transfers a specified amount from one account to another. Returns an error if the withdrawal or deposit fails.

## Usage Example

``` go
package main

import (
    "fmt"
    "path/to/bank"
)

func main() {
    // Create customers
    customer1 := bank.Customer{Name: "Alice", Address: "123 Main St", Phone: "123-456-7890"}
    customer2 := bank.Customer{Name: "Bob", Address: "456 Elm St", Phone: "987-654-3210"}

    // Create accounts
    account1 := &bank.Account{Customer: customer1, Number: 1001, Balance: 500.0}
    account2 := &bank.Account{Customer: customer2, Number: 1002, Balance: 300.0}

    // Deposit
    err := account1.Deposit(200)
    if err != nil {
        fmt.Println("Deposit error:", err)
    }

    // Withdraw
    err = account2.Withdraw(100)
    if err != nil {
        fmt.Println("Withdraw error:", err)
    }

    // Transfer
    err = account1.Transfer(account2, 150)
    if err != nil {
        fmt.Println("Transfer error:", err)
    }

    // Print statements
    fmt.Println(account1.Statement())
    fmt.Println(account2.Statement())
}
```

## Error Handling

The methods in the Account type return errors in the following scenarios:

- **Deposit:** If the amount to be deposited is less than or equal to zero.
- **Withdraw:** If the amount to be withdrawn is greater than the balance or less than or equal to zero.
- **Transfer:** If either the withdrawal from the source account or the deposit to the destination account fails.

## Testing

The package includes a test file bank_test.go with tests for the Account type methods.
Running Tests

To run the tests, use the go test command:
``` sh
go test path/to/bank
```

### Tests Included

- **TestAccount:** Verifies that an Account object can be created successfully.
- **TestDeposit:** Ensures that the balance is updated correctly after a deposit.
- **TestDepositNegative:** Verifies that depositing a negative amount returns an error.
- **TestWithdraw:** Ensures that the balance is updated correctly after a withdrawal.
- **TestWithdrawNegative:** Verifies that withdrawing a negative amount returns an error.
- **TestStatement:** Checks that the account statement is generated correctly.
- **TestAccountTransfer:** Tests the transfer of funds between accounts.

### Example Test

Here is an example test for the Deposit method:

``` go
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
```

## Contributing

Contributions to the bank package are welcome. Please submit a pull request or open an issue to discuss any changes or improvements.

## License

This project is licensed under the MIT License. See the LICENSE file for details.