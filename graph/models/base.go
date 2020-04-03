package models

import (
	"github.com/gofrs/uuid"
	"log"
	"time"
)

type BaseModel struct {
	ID 			string 		`gorm:"primary_key;";json:"id"`
	CreatedAt   time.Time 	`json:"createdAt"`
	UpdatedAt   time.Time 	`gorm:"index";json:"updatedAt"`
	DeletedAt   *time.Time 	`gorm:"index"`
}

type BaseChangeModel struct {
	CreatedByID *string 	`json:"createdById"`
	UpdatedByID *string 	`json:"updatedById"`
	DeletedByID *string 	`json:"deletedById"`
}

func (bm *BaseModel) BeforeSave() (err error) {
	u, err := uuid.NewV4()
	if err != nil {
		log.Printf("Failed to generate UUID: %v\n", err)
		return err
	}
	bm.ID = u.String()
	return nil
}
