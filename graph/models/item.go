package models

import "time"

type Item struct {
	ID        int    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string `json:"name"`
	DeletedAt *time.Time `json:"updatedAt"`
}
