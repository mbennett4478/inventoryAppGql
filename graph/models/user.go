package models

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	FirstName 	string		`json:"firstName"`
	LastName 	string		`json:"lastName"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	Roles		[]Role		`gorm:"many2many:user_roles"`
}

type Role struct {
	BaseModel
	BaseChangeModel
	Name 	string 	`json:"name"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uu, err := uuid.NewV4()
	if err != nil {
		return err
	}
	u.ID = uu.String()

	if u.Password != "" {
		if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 11); err == nil {
			_ = scope.SetColumn("Password", pw)
		}
	}
	return nil
}

