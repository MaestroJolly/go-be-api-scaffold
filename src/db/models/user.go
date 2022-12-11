package models

import (
	"html"
	"strings"

	database "github.com/MaestroJolly/go-be-api-scaffold/src/db"
	"github.com/MaestroJolly/go-be-api-scaffold/src/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	FirstName string `gorm:"size:255;not null;" json:"firstname"`
	LastName  string `gorm:"size:255;not null;" json:"lastname"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Password  string `gorm:"size:255;not null;" json:"-"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := helpers.Encryptor(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(*passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
