package types

import "time"

type User struct {
	ID        string    `json:"id" binding:"omitempty,uuid"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required,min=8"`
	CreatedAt time.Time `json:"createdAt"`
}
