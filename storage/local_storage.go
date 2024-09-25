package storage

import (
	"errors"
	"fmt"
)

// Used to store history in memory for unit tests
type localStorage struct {
	history []HistoryEntry
	users   map[string]*User
}

func NewLocalStorage() *localStorage {
	return &localStorage{
		history: []HistoryEntry{},
		users:   make(map[string]*User),
	}
}

// Save the history entry to the localStorage
func (storage *localStorage) SaveOperation(entry HistoryEntry) error {
	storage.history = append(storage.history, entry)
	return nil
}

// Get the history from the localStorage
func (storage *localStorage) GetHistory() ([]HistoryEntry, error) {
	if len(storage.history) == 0 {
		return nil, errors.New("no history found")
	}
	return storage.history, nil
}

// Reset the history in the localStorage
func (storage *localStorage) ResetHistory() error {
	storage.history = []HistoryEntry{}
	return nil
}

func (storage *localStorage) RegisterUser(username string, password string) error {

	if storage.users[username].Username != "" {
		return errors.New("user already exists")
	}

	newUser := NewUser(username, password)

	storage.users[username] = newUser
	return nil
}

func (userStorage *localStorage) AuthenticateUser(username string, password string) error {

	user := userStorage.users[username]

	if user == nil {
		return fmt.Errorf("user %s not found", username)
	}

	if user.Password != password {
		return fmt.Errorf("wrong password for username: %s", username)
	}

	return nil
}
