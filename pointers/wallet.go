package pointers

import (
	"errors"
	"fmt"
)

// a new type we can declare methods on
type Bitcoin int

// in Go if a symbol starts with lowercase then it is private outside the package it is defined in
type Wallet struct {
	balance Bitcoin
}

// lets you define how your type is printed when used with the %s format string in prints
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
