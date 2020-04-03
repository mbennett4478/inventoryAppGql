package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Item struct {
	BaseModel
	BaseChangeModel
	Name      string `json:"name"`
}

func (i *Item) BeforeCreate(scope *gorm.Scope) error {
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}
	err = scope.SetColumn("ID", u.String())
	return err
}
