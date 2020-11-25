// Package factory example pattern.
package factory

import (
	"fmt"
)

// Const types payments.
const (
	Cash int = iota + 1
	DebitCard
)

// PaymentMethod ...
type PaymentMethod interface {
	Pay(amount float32) string
}

// GetPaymentMethod returning payment.
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	}

	return nil, fmt.Errorf("payment method %d not recognized", m) //nolint:goerr113
}

// CashPM ...
type CashPM struct{}

// Pay ...
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using cash\n", amount)
}

// DebitCardPM ...
type DebitCardPM struct{}

// Pay ...
func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}
