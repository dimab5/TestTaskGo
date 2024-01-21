package walletActions

import (
	"TestTaskGo/internal/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type ResponseGetWallet struct {
	ID      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}

type GetWallet interface {
	FindWallet(id uuid.UUID) (models.Wallet, error)
}

func NewGetWallet(getWallet GetWallet) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID, err := uuid.Parse(chi.URLParam(r, "walletId"))
		if err != nil {
			http.Error(w, "Invalid wallet ID", http.StatusNotFound)
			return
		}

		wallet, err := getWallet.FindWallet(walletID)
		if err != nil {
			http.Error(w, "Wallet not found", http.StatusNotFound)
			return
		}

		response := ResponseGetWallet{
			ID:      wallet.ID,
			Balance: wallet.Balance,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
