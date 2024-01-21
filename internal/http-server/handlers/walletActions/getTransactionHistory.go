package walletActions

import (
	"TestTaskGo/internal/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ResponseGetTransactionHistory struct {
	FromID         uuid.UUID `json:"from_id"`
	ToID           uuid.UUID `json:"to_id"`
	Amount         float64   `json:"amount"`
	Time           time.Time `json:"time"`
	ResponseStatus int       `json:"response_status"`
}

type TransactionHistory interface {
	GetTransactionsById(id uuid.UUID) ([]models.Transaction, error)

	FindWallet(id uuid.UUID) (models.Wallet, error)
}

func NewTransactionHistory(transactionHistory TransactionHistory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID, err := uuid.Parse(chi.URLParam(r, "walletId"))
		if err != nil {
			http.Error(w, "Invalid wallet ID", http.StatusNotFound)
			return
		}

		transactions, err := transactionHistory.GetTransactionsById(walletID)
		_, err2 := transactionHistory.FindWallet(walletID)
		if err != nil || err2 != nil {
			http.Error(w, "Wallet not found", http.StatusNotFound)
			return
		}

		var response []ResponseGetTransactionHistory
		for _, t := range transactions {
			response = append(response, ResponseGetTransactionHistory{
				FromID:         t.FromID,
				ToID:           t.ToID,
				Amount:         t.Amount,
				Time:           t.Time,
				ResponseStatus: t.ResponseStatus,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
