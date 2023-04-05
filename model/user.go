package model

import (
	"log"
	"time"

	"github.com/fuadop/rapptr-task/types"
)

func (m *Model) InsertUser(username, password string) error {
	query := `
		INSERT INTO users(id, username, password, createdAt)
		VALUES(UUID(), ?, MD5(?), NOW());
	`
	if _, err := m.DB.Exec(query, username, password); err != nil {
		return err
	}

	return nil
}

func (m *Model) GetUser(id string) (*types.User, error) {
	query := "SELECT username, password, createdAt FROM users WHERE id = ?;"

	var createdAt time.Time
	var username, password string
	if err := m.DB.QueryRow(query, id).Scan(&username, &password, &createdAt); err != nil {
		return nil, err
	}

	user := &types.User{ID: id, Username: username, Password: password, CreatedAt: createdAt}
	return user, nil
}

func (m *Model) GetUserByUsername(username string) (*types.User, error) {
	query := "SELECT id, password, createdAt FROM users WHERE username = ?;"

	var createdAt time.Time
	var id, password string
	if err := m.DB.QueryRow(query, username).Scan(&id, &password, &createdAt); err != nil {
		return nil, err
	}

	user := &types.User{ID: id, Username: username, Password: password, CreatedAt: createdAt}
	return user, nil
}

func (m *Model) HashPassword(password string) string {
	query := "SELECT MD5(?);"

	var hashedPassword string
	if err := m.DB.QueryRow(query, password).Scan(&hashedPassword); err != nil {
		log.Println(err)
	}

	return hashedPassword
}
