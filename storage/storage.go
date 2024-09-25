package storage

import (
	"time"
)

// HistoryEntry represents a calculator operation in history.
type HistoryEntry struct {
	Operand1  float64   // Left operand in expression
	Operand2  float64   // Right operand in expression
	Operation string    // +, -, *, /, %, ^
	Result    float64   // Result after applying the operation
	Timestamp time.Time // When the operation was performed
}

type User struct {
	Username string
	Password string
}

func NewUser(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

type Storage interface {
	// Operations methods related to calculator history
	SaveOperation(entry HistoryEntry) error
	GetHistory() ([]HistoryEntry, error)
	ResetHistory() error

	// User related methods
	RegisterUser(username string, password string) error
	AuthenticateUser(username string, password string) error
}
