package models

import (
	"github.com/gofrs/uuid"
)

type InventoryItem struct {
	BaseModelSoftDelete
	InventoryID uuid.UUID	`json:"inventoryId"`
	Quantity    int    		`json:"quantity"`
	Item        Item  		`json:"item"`
	ItemID		int			`json:"-"`
}
