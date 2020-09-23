package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is
type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is
type Wallet struct {
	balance Bitcoin
}

// Deposit is
func (w *Wallet) Deposit(cash Bitcoin) {
	w.balance += cash
}

// ErrInsufficientBalance is
var ErrInsufficientBalance = errors.New("ops, insufficient balance")

// Withdrawn is
func (w *Wallet) Withdrawn(cash Bitcoin) error {
	if cash > w.balance {
		return ErrInsufficientBalance
	}
	w.balance -= cash

	return nil
}

// Balance is
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
