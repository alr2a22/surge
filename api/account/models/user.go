package models

import (
	"encoding/json"
	"surge/internal/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	db.Model
	Username  string         `json:"username" gorm:"unique" validate:"required"`
	Password  string         `json:"password,omitempty"`
	LastLogin gorm.DeletedAt `json:"-"`
	Admin     bool           `json:"-" gorm:"default:0"`
}

func (user *User) SetPassword(rawPassword string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.MinCost)
	user.Password = string(hash)
}

func (user *User) CheckPassword(rawPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rawPassword))
}

func (user User) MarshalJSON() ([]byte, error) {
	type u User // prevent recursion
	x := u(user)
	x.Password = ""
	return json.Marshal(x)
}
