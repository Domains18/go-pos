package domain

import "time"

type PaymentType string

const (
	cash PaymentType = "CASH"
	EWallet PaymentType = "E-Wallet"
	EDC PaymentType = "EDC"
)

type Payment struct {
	ID        uint64      `json:"id"`
	Name      string      `json:"name"`
	Type      PaymentType `json:"type"`
	Logo      string      `json:"logo"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
