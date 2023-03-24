package entity

import uuid "github.com/satori/go.uuid"

type User struct {
	Id       uuid.UUID `json:"-" db:"id"`
	Name     string    `json:"name" binding:"required"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
}
