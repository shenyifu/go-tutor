package point

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var MoneyNotEnoughError = errors.New("out of balance")

func (w *Wallet) Deposit(incom Bitcoin) {
	w.balance += incom
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return MoneyNotEnoughError
	}
	w.balance -= amount
	return nil
}
