package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type BaseModel struct {
	ID 			uuid.UUID 	`gorm:"primary_key;type:uuid;";json:"id"`
	CreatedByID *uuid.UUID 	`gorm:"type:uuid";json:"createdById"`
	UpdatedByID *uuid.UUID 	`gorm:"type:uuid";json:"updatedById"`
	CreatedAt   time.Time 	`json:"createdAt"`
	UpdatedAt   time.Time 	`gorm:"index";json:"updatedAt"`
}

type BaseModelSoftDelete struct {
	BaseModel
	DeletedByID *uuid.UUID `gorm:"type:uuid"`
	DeletedAt   *time.Time `gorm:"index"`
}
