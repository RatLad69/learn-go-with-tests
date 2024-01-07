package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Coin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted error. Did not get one")
		}

		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got error. Did not want one")
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Coin(10))
		assertBalance(t, wallet, Coin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Coin(20)}
		err := wallet.Withdraw(Coin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Coin(10))
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		startingBalance := Coin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Coin(25))

		assertError(t, err, ErrInsufficient)
		assertBalance(t, wallet, startingBalance)
	})
}
