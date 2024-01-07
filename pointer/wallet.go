package main

import (
	"errors"
	"fmt"
)

type Coin int

var ErrInsufficient = errors.New("cannot withdraw, insufficient funds")

func (c Coin) String() string {
	return fmt.Sprintf("%d Coins", c)
}

type Wallet struct {
	balance Coin
}

func (w *Wallet) Deposit(amount Coin) {
	fmt.Printf("Address of balance in Deposit is %p\n", &w.balance)
	(*w).balance += amount
}

func (w *Wallet) Withdraw(amount Coin) error {
	if amount > w.balance {
		return ErrInsufficient
	}

	(*w).balance -= amount
	return nil
}

func (w *Wallet) Balance() Coin {
	return (*w).balance
}
