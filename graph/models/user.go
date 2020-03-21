package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModelSoftDelete
	FirstName 	string		`json:"firstName"`
	LastName 	string		`json:"lastName"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}

func (u *User) BeforeSave(scope *gorm.Scope) error {
	if u.Password != "" {
		if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 11); err != nil {
			_ = scope.SetColumn("Password", pw)
		}
	}
	return nil
}
