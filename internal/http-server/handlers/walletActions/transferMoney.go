package walletActions

import (
	"TestTaskGo/internal/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type RequestTransferMoney struct {
	ToWalletId uuid.UUID `json:"to"`
	Amount     float64   `json:"amount"`
}

type ResponseTransferMoney struct {
	Status int `json:"status"`
}

type StorageInterface interface {
	MoneyTransfer(
		from_ID,
		to_ID uuid.UUID,
		amount float64,
		response_status int) error

	FindWallet(id uuid.UUID) (models.Wallet, error)

	ChangeBalance(id uuid.UUID, newBalance float64) error
}

func NewTransferMoney(storageInterface StorageInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request RequestTransferMoney

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Error decoding JSON request", http.StatusBadRequest)
			return
		}

		walletID, err := uuid.Parse(chi.URLParam(r, "walletId"))
		if err != nil {
			http.Error(w, "Wallet not found", http.StatusNotFound)
			return
		}

		toWallet, err := storageInterface.FindWallet(request.ToWalletId)
		fromWallet, _ := storageInterface.FindWallet(walletID)
		if err != nil || fromWallet.Balance < request.Amount {
			http.Error(w, "Wallet not found or not enough money", http.StatusBadRequest)
			return
		}

		response := ResponseTransferMoney{Status: http.StatusOK}
		storageInterface.MoneyTransfer(
			walletID,
			request.ToWalletId,
			request.Amount,
			response.Status,
		)

		storageInterface.ChangeBalance(walletID, fromWallet.Balance-request.Amount)
		storageInterface.ChangeBalance(request.ToWalletId, toWallet.Balance+request.Amount)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
