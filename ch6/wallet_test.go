package ch6_test

import (
	"learn-go-with-tests/ch6"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := ch6.Wallet{}
	wallet.Deposit(ch6.Bitcoin(10))

	got := wallet.Balance()

	want := ch6.Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
