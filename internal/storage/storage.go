package storage

import (
	"TestTaskGo/internal/config"
	"TestTaskGo/internal/models"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Storage struct {
	Db *sql.DB
}

func ConnectToDB() (*sql.DB, error) {
	cfg := config.MustLoad()
	connStr := cfg.DatabaseConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *Storage) AddWallet() (uuid.UUID, error) {
	newUUID := uuid.New()

	query := "INSERT INTO Wallets (id, balance) VALUES ($1, $2)"

	_, err := s.Db.Exec(query, newUUID, 100.0)
	return newUUID, err
}

func (s *Storage) MoneyTransfer(
	from_ID,
	to_ID uuid.UUID,
	amount float64,
	response_status int) error {
	query := `
		INSERT INTO Transaction_History (from_id, to_id, amount, time, response_status)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := s.Db.Exec(query, from_ID, to_ID, amount, time.Now().Format(time.RFC3339), response_status)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) FindWallet(id uuid.UUID) (models.Wallet, error) {
	query := "SELECT id, balance FROM Wallets WHERE id = $1"
	row := s.Db.QueryRow(query, id)

	var wallet models.Wallet
	err := row.Scan(&wallet.ID, &wallet.Balance)
	if err != nil {
		return models.Wallet{}, err
	}

	return wallet, nil
}

func (s *Storage) ChangeBalance(id uuid.UUID, newBalance float64) error {
	query := "UPDATE Wallets SET balance = $1 WHERE id = $2"
	_, err := s.Db.Exec(query, newBalance, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetTransactionsById(id uuid.UUID) ([]models.Transaction, error) {
	rows, err := s.Db.Query(`
		SELECT from_id, to_id, amount, time, response_status 
		FROM Transaction_History
		WHERE from_id = $1 OR to_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.FromID, &t.ToID, &t.Amount, &t.Time, &t.ResponseStatus)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
