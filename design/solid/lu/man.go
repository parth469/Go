package main

import "fmt"

// Step 1: Define separate interfaces
type DigitalPayment interface {
	Pay(amount float64) string
}

type CardPayment interface {
	UseCard() string
}

// Step 2: Implement CreditCard payment (supports both interfaces)
type CreditCard struct{}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card", amount)
}

func (c CreditCard) UseCard() string {
	return "Using Credit Card"
}

// Step 3: Implement Wallet Payment (only implements DigitalPayment)
type WalletPayment struct{}

func (w WalletPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Wallet", amount)
}

// Step 4: Functions handling different payments
func ProcessDigitalPayment(p DigitalPayment, amount float64) {
	fmt.Println(p.Pay(amount))
}

func ProcessCardPayment(cp CardPayment) {
	fmt.Println(cp.UseCard())
}

func main() {
	creditCard := CreditCard{}
	wallet := WalletPayment{}

	ProcessDigitalPayment(creditCard, 100.50) // ✅ Works fine
	ProcessDigitalPayment(wallet, 50.00)      // ✅ Works fine
	ProcessCardPayment(creditCard)            // ✅ Works fine
	// ProcessCardPayment(wallet) // ❌ Compile-time error (Wallet does not support card payments)
}
