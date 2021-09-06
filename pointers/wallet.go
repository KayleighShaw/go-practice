package pointers

import "fmt"

// a new type we can declare methods on
type Bitcoin int

// in Go if a symbol starts with lowercase then it is private outside the package it is defined in
type Wallet struct {
	balance Bitcoin
}

// Pointers let us point to some values and let us change them
// Rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so we can change the original values
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
