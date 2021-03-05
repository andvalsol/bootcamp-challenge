package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Entry point
/*
As a user I have to be able to:

* Deposit and withdraw funds from my accounts
* Get my account balance
* Get my historical transactions (only the amounts)
* Transfer funds to other accounts (can be to a distinct user)

Validations:

* Can't withdraw non-existing funds from your account.
* Can't make transactions greater than $10,000 in a single exhibition.
*/

const maxTransactionAmount = 10000

var historicalTransactions = make(map[string][]float64)

func main() {
	// Create a default account
	defaultAccount := Account{
		ID:      "Acc1",
		Balance: 0,
	}

	user1 := generateUser("Andrey", "ID1", append([]Account{}, defaultAccount))
	user2 := generateUser("Juan", "ID2", append([]Account{}, defaultAccount))

	fmt.Println("User 1 -> ", user1)
	fmt.Println("User 2 -> ", user2)

	// deposit $5000 to user1 account
	err := depositToAccount(&user1.Accounts[0], 5000)

	if err == nil {
		result := fmt.Sprintf("the amount of money of the user1 is: %v", user1.GetBalance())
		fmt.Print(result)
	} else {
		fmt.Print(err.Error())
	}

	// Transfer $5,000 from User 1 Account 1 to User 1 Account 2
	Transfer(&user1.Accounts[0], &user2.Accounts[0], 5000)
	result := fmt.Sprintf("the amount of money of the user1 is: %v", user1.GetBalance())
	fmt.Print(result)

	// Transfer $15,000 from User 1 Account 2 to User 2 Account 1 (can use multiple transactions)
	fmt.Println()

	// Print current balance of User 1 Account 2 & User 2 Account 1
	fmt.Println()

	// Print historic deposits & withdrawals for each user
	for _, transactions := range historicalTransactions {
		for _, transaction := range transactions {
			fmt.Print("Transaction is: " + strconv.FormatFloat(transaction, 'E', -1, 64))
		}
	}
	fmt.Println()

	// Try to transfer $20,000 from User 1 Account 1 to User 2 Account 1
	err2 := Transfer(&user1.Accounts[0], &user2.Accounts[0], 20000)
	if err2 == nil {
		fmt.Println(err2.Error())
	}
	fmt.Println()

	// Try to transfer $11,000 from User 2 Account 1 to User 1 Account 1
	fmt.Println()
}

type LimitExceededError struct {
	Err error
}

func (error *LimitExceededError) Error() string {
	return fmt.Sprintf("limit exceeded")
}

func (error *LimitExceededError) Unwrap() error { return error.Err }

func depositToAccount(account *Account, amount float64) error {
	// Check that the amount doesn't exceed the limit
	if amount > maxTransactionAmount {
		return &LimitExceededError{
			Err: errors.New("limit exceeded"),
		}
	}

	// Proceed to deposit the money into the proper account
	account.Balance += amount

	return nil
}

func generateUser(name string, id string, accounts []Account) User {
	return User{
		Name:     name,
		ID:       id,
		Accounts: accounts,
	}
}

// Create the User type
type User struct {
	Name     string
	ID       string
	Accounts []Account
}

func (user *User) GetBalance() float64 {
	var balance float64 = 0

	for _, account := range user.Accounts {
		balance += account.Balance
	}

	return balance
}

func Transfer(from *Account, to *Account, amountToTransfer float64) error {
	// Check if the transaction exceeds the amount to transfer
	if amountToTransfer > maxTransactionAmount {
		return &LimitExceededError{
			Err: errors.New("limit exceeded"),
		}
	}

	from.Balance -= amountToTransfer
	to.Balance += amountToTransfer

	// Add the transaction to the list of transactions
	// In case of the account that the money is being reduced the transaction is negative
	historicalTransactions[from.ID] = append(historicalTransactions[from.ID], -amountToTransfer)

	// In case of the account that the money is being added the transaction is positive
	historicalTransactions[to.ID] = append(historicalTransactions[to.ID], amountToTransfer)

	return nil
}

type Account struct {
	ID      string
	Balance float64
}

type Transaction struct {
	ID          string
	Amount      float64
	Origin      Account
	Destination Account
	Date        time.Time
}
