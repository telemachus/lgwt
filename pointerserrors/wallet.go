package pointerserrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("you can’t withdraw more than you have!")

func (w *Wallet) Deposit(amount Bitcoin) Bitcoin {
	w.balance += amount
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) (Bitcoin, error) {
	if amount > w.balance {
		return w.balance, ErrInsufficientFunds
	}

	w.balance -= amount
	return w.balance, nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
