package model

import (
	"log"
	"time"

	"github.com/fuadop/rapptr-task/types"
)

func (m *Model) InsertTodo(text, userId string) error {
	query := `
	INSERT INTO todos(id, userId, text, createdAt, updatedAt)
	VALUES(UUID(), ?, ?, NOW(), NOW());
	`
	if _, err := m.DB.Exec(query, userId, text); err != nil {
		return err
	}

	return nil
}

func (m *Model) UpdateTodo(id, text string) error {
	query := `
		UPDATE todos
		SET text = ?, updatedAt = NOW()
		WHERE id = ?;
	`
	if _, err := m.DB.Exec(query, text, id); err != nil {
		return err
	}

	return nil
}

func (m *Model) GetTodos(userId string) ([]types.Todo, error) {
	todos := make([]types.Todo, 0)

	query := `
		SELECT id, text, createdAt, updatedAt
		FROM todos
		WHERE userId = ? AND isDeleted IS NOT TRUE;
	`
	rows, err := m.DB.Query(query, userId)
	if err != nil {
		return todos, err
	}

	var id, text string
	var createdAt, updatedAt time.Time
	for rows.Next() {
		if err = rows.Scan(&id, &text, &createdAt, &updatedAt); err != nil {
			log.Println(err)
		}

		todo := types.Todo{ID: id, UserID: userId, Text: text, CreatedAt: createdAt, UpdatedAt: updatedAt}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (m *Model) GetTodo(id string) (*types.Todo, error) {
	query := `
		SELECT userId, text, createdAt, updatedAt
		FROM todos 
		WHERE id = ? AND isDeleted IS NOT TRUE;
	`

	var userId, text string
	var createdAt, updatedAt time.Time
	if err := m.DB.QueryRow(query, id).Scan(&userId, &text, &createdAt, &updatedAt); err != nil {
		return nil, err
	}

	todo := &types.Todo{ID: id, UserID: userId, Text: text, CreatedAt: createdAt, UpdatedAt: updatedAt}
	return todo, nil
}

func (m *Model) DeleteTodo(id string) error {
	query := `
		UPDATE todos
		SET isDeleted = TRUE
		WHERE id = ?;
	`
	if _, err := m.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}
