package models

import (
	"html"
	"strings"
	"time"

	database "github.com/MaestroJolly/go-be-api-scaffold/src/db"
	"github.com/MaestroJolly/go-be-api-scaffold/src/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`
	Username  string         `gorm:"size:255;not null;unique" json:"username"`
	FirstName string         `gorm:"size:255;not null;" json:"first_name"`
	LastName  string         `gorm:"size:255;not null;" json:"last_name"`
	Email     string         `gorm:"size:255;not null;unique" json:"email"`
	Password  string         `gorm:"size:255;not null;" json:"-"`
	CreatedAt time.Time      `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
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

func FindUserByUsernameOrEmail(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Or("email", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
