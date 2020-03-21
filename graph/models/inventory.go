package models

type Inventory struct {
	BaseModelSoftDelete
	Name      	string 			`json:"name"`
	Items     	[]InventoryItem `json:"items"`
}
