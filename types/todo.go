package types

import "time"

type Todo struct {
	ID        string    `json:"id" binding:"omitempty,uuid"`
	UserID    string    `json:"-" binding:"omitempty,uuid"`
	Text      string    `json:"text" binding:"required,min=1"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	IsDeleted bool `json:"-"` // for soft delete
}
