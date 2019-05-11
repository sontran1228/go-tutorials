package pointers

import "fmt"

// Wallet has balance information of a wallet
type Wallet struct {
	// Spec: In Go if a symbol (so variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
	// balance is package private because it starts with lowercase symbol "b"
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
