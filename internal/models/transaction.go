package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	FromID         uuid.UUID `json:"from_id"`
	ToID           uuid.UUID `json:"to_id"`
	Amount         float64   `json:"amount"`
	Time           time.Time `json:"time"`
	ResponseStatus int       `json:"response_status"`
}
