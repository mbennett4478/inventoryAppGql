package models

type InventoryItem struct {
	BaseChangeModel
	InventoryID string	`gorm:"primary_key";json:"inventoryId"`
	Quantity    int    	`json:"quantity"`
	Item        Item  	`json:"item"`
	ItemID		string 	`gorm:"primary_key";json:"-"`
}
