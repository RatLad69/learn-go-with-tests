package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	fmt.Printf("Address of balance in test is %p\n", &wallet.balance)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}
