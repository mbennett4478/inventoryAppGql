package models

import "time"

type Inventory struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Items     []InventoryItem `json:"items"`
}
