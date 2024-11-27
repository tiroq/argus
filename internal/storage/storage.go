package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteUserStorage struct {
	db *sql.DB
}

func NewSQLiteUserStorage(dbPath string) (*SQLiteUserStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	storage := &SQLiteUserStorage{db: db}
	if err := storage.initDB(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *SQLiteUserStorage) initDB() error {
	query := `
    CREATE TABLE IF NOT EXISTS subscriptions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER UNIQUE
    );`
	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	return nil
}

func (s *SQLiteUserStorage) AddUserSubscription(userID int) error {
	query := `INSERT INTO subscriptions (user_id) VALUES (?) ON CONFLICT(user_id) DO NOTHING;`
	_, err := s.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to add user subscription: %w", err)
	}
	return nil
}

func (s *SQLiteUserStorage) IsUserSubscribed(userID int) (bool, error) {
	query := `SELECT user_id FROM subscriptions WHERE user_id = ? LIMIT 1;`
	row := s.db.QueryRow(query, userID)

	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (s *SQLiteUserStorage) GetAllSubscribedUsers() ([]int, error) {
	query := `SELECT user_id FROM subscriptions;`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch subscribed users: %w", err)
	}
	defer rows.Close()

	var userIDs []int
	for rows.Next() {
		var userID int
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}

	return userIDs, nil
}
