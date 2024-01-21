package tests

import (
	"TestTaskGo/internal/http-server/handlers/walletActions"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockWalletCreator struct {
	mock.Mock
}

func (m *MockWalletCreator) AddWallet() (uuid.UUID, error) {
	args := m.Called()
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func TestNewWalletCreatorHandler(t *testing.T) {
	mockWalletCreator := new(MockWalletCreator)
	mockGeneratedID := uuid.New()

	mockWalletCreator.On("AddWallet").Return(mockGeneratedID, nil)

	handler := walletActions.NewWalletCreator(mockWalletCreator)

	request := httptest.NewRequest("POST", "/api/v1/wallet", bytes.NewReader([]byte("")))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	var response walletActions.ResponseCreateWallet
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, mockGeneratedID, response.ID)
	assert.Equal(t, float64(100.0), response.Balance)
	assert.Equal(t, http.StatusOK, response.Status)
	assert.Empty(t, response.Error)

	mockWalletCreator.AssertExpectations(t)
}
