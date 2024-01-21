package models

import "github.com/google/uuid"

type Wallet struct {
	ID      uuid.UUID `db:"id"`
	Balance float64   `db:"balance"`
}
