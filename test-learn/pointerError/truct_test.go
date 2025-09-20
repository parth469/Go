package pointererror

import (
	"math/rand"
	"testing"
)

func TestWallet(t *testing.T) {
	initialBalance := rand.Intn(100) + 20 // Ensure enough balance for withdrawal tests
	w := Wallet{balance: initialBalance}

	checkBalance := func(t *testing.T, want int) {
		got := w.CheckBalance()
		if got != want {
			t.Errorf("CheckBalance() = %d; want %d", got, want)
		}
	}

	t.Run("Check Balance", func(t *testing.T) {
		checkBalance(t, initialBalance)
	})

	t.Run("Deposit", func(t *testing.T) {
		depositAmount := 10
		err := w.Deposit(depositAmount)
		if err != nil {
			t.Errorf("Deposit() error = %v; want nil", err)
		}
		checkBalance(t, initialBalance+depositAmount)
	})

	t.Run("Withdrawal with sufficient balance", func(t *testing.T) {
		withdrawAmount := 10
		currentBalance := w.CheckBalance() // Get current balance before withdrawal
		err := w.Withdrawal(withdrawAmount)
		if err != nil {
			t.Errorf("Withdrawal() error = %v; want nil", err)
		}
		checkBalance(t, currentBalance-withdrawAmount)
	})

	t.Run("Withdrawal with insufficient balance", func(t *testing.T) {
		currentBalance := w.CheckBalance() // Get current balance before attempting withdrawal
		overdrawAmount := currentBalance + 50
		err := w.Withdrawal(overdrawAmount)
		if err == nil {
			t.Errorf("Withdrawal() error = nil; want error due to insufficient balance")
		}
		checkBalance(t, currentBalance)
	})
}