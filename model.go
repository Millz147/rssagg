package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/millz147/rssagg/internal/database"
)

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func convertDbUserToUser(user database.User) User{

return User{
	ID: user.ID,
	Name: user.Name,
	CreatedAt: user.CreatedAt,
	UpdatedAt: user.UpdatedAt,
}

}