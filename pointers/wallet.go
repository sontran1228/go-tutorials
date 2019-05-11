package pointers

import "fmt"

// Stringer is used to format the data to be a string
type Stringer interface {
	String() string
}

// Bitcoin is a declaration of new type from existing ones which is "int"
type Bitcoin int

// by making a new type, we can declare methods on them
// to check the result of this formater method, change the test case to make it failed.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet has balance information of a wallet
type Wallet struct {
	// Spec: In Go if a symbol (so variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
	// balance is package private because it starts with lowercase symbol "b"
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
