package walletActions

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type ResponseCreateWallet struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Balance float64   `json:"balance,omitempty"`
	Status  int       `json:"status"`
	Error   string    `json:"error,omitempty"`
}

type WalletCreator interface {
	AddWallet() (uuid.UUID, error)
}

func NewWalletCreator(walletCreator WalletCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		generatedID, err := walletCreator.AddWallet()
		if err != nil {
			response := ResponseCreateWallet{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}
			sendJSONResponse(w, response, http.StatusBadRequest)
			return
		}

		response := ResponseCreateWallet{
			ID:      generatedID,
			Balance: 100.0,
			Status:  http.StatusOK,
		}
		sendJSONResponse(w, response, http.StatusOK)
	}
}

func sendJSONResponse(w http.ResponseWriter, response ResponseCreateWallet, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
