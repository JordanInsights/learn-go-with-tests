package ch6_test

import (
	"learn-go-with-tests/ch6"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet ch6.Wallet, want ch6.Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := ch6.Wallet{}
		wallet.Deposit(ch6.Bitcoin(10))
		want := ch6.Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := ch6.Wallet{}
		wallet.Deposit(10)
		wallet.Withdraw(5)

		want := ch6.Bitcoin(5)
		assertBalance(t, wallet, want)
	})
}
