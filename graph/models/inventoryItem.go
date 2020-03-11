package models

type InventoryItem struct {
	ID          int    	`json:"id"`
	InventoryID int    	`json:"inventoryId"`
	Quantity    int    	`json:"quantity"`
	Item        Item  	`json:"item"`
	ItemID		int		`json:"-"`
}
