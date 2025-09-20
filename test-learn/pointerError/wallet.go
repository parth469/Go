package pointererror

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) error {
	w.balance += amount

	return nil
}

func (w Wallet) CheckBalance() int {
	return w.balance
}

func (w *Wallet) Withdrawal(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("insufficient balance: attempted to withdraw %d, but current balance is %d", amount, w.balance)
	}

	w.balance -= amount

	return nil
}
