package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("should return the wallet balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("should return the amount withdraw from the wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdrawn(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
		checkInexistErr(t, err)
	})

	t.Run("should return the insufficient balance message", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		err := wallet.Withdrawn(Bitcoin(100))

		checkBalance(t, wallet, initialBalance)
		checkErr(t, err, ErrInsufficientBalance)
	})
}

func checkBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("\ngot: %s\nwant: %s", got, want)
	}
}

func checkErr(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error, but none occurred")
	}

	if got != want {
		t.Errorf("\ngot: %s\nwant: %s", got, want)
	}
}

func checkInexistErr(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("unexpected error")
	}
}
