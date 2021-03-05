package main

/*
import (
	"fmt"
	"time"
)

const TransferLimit = 10000

type User struct {
	ID uint
	Name string
	Accounts []Account
}

type Account struct {
	ID string
	Balance float64
}

type Transaction struct {
	ID uint
	Amount float64
	Origin *Account
	Destiny *Account
	Date time.Time
}

// Here we save the amount of the deposits & withdrawals for each Account
var historical map[string][]float64

func main (){
	user1,user2 := generateUsers()
	fmt.Println("User 1 -> ",user1)
	fmt.Println("User 2 -> ",user2)
	fmt.Println()
	// Print current balance of all the accounts from user 1 & user 2
	fmt.Println()
	// Transfer $5,000 from User 1 Account 1 to User 1 Account 2
	fmt.Println()
	// Print current balance of User 1 Account 1 & User 1 Account 2
	fmt.Println()
	// Transfer $15,000 from User 1 Account 2 to User 2 Account 1 (can use multiple transactions)
	fmt.Println()
	// Print current balance of User 1 Account 2 & User 2 Account 1
	fmt.Println()
	// Print historic deposits & withdrawals for each user
	fmt.Println()
	// Try to transfer $20,000 from User 1 Account 1 to User 2 Account 1
	fmt.Println()
	// Try to transfer $11,000 from User 2 Account 1 to User 1 Account 1
	fmt.Println()
}

func GetBalance(account Account){
	// CODE HERE
}

func GetHistoricalTransactions(user User){
	// CODE HERE
}

func MakeTransaction(transaction Transaction) error{
	//CODE HERE
	return nil
}

func generateUsers() (user1,user2 User){
	// USER 1
	user1 = User{
		ID:       1,
		Name:     "User #1",
	}
	account1User1 := Account{
		ID:      "1111-1111-1111-1111",
		Balance: 20000,
	}
	account2User1 := Account{
		ID:     "2222-2222-2222-2222",
		Balance: 20000,
	}
	user1.Accounts = append(user1.Accounts,account1User1)
	user1.Accounts = append(user1.Accounts,account2User1)
	// USER 2
	user2 = User{
		ID:       2,
		Name:     "User #2",
	}
	account1User2 := Account{
		ID:      "3333-3333-3333-3333",
		Balance: 20000,
	}
	account2User2 := Account{
		ID:     "4444-4444-4444-4444",
		Balance: 20000,
	}
	user2.Accounts = append(user2.Accounts,account1User2)
	user2.Accounts = append(user2.Accounts,account2User2)
	return
}
*/
