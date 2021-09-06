package pointers

import "fmt"

// in Go if a symbol starts with lowercase then it is private outside the package it is defined in
type Wallet struct {
	balance int
}

// Pointers let us point to some values and let us change them
// Rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so we can change the original values
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
