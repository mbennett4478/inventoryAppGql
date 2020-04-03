package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Inventory struct {
	BaseModel
	BaseChangeModel
	Name	string 			`json:"name"`
	Items 	[]InventoryItem `json:"items"`
}

func (i* Inventory) BeforeCreate(scope *gorm.Scope) error {
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}
	err = scope.SetColumn("ID", u.String())
	return err
}
