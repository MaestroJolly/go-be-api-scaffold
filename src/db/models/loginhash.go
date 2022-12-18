package models

import (
	"time"

	database "github.com/MaestroJolly/go-be-api-scaffold/src/db"
	"gorm.io/gorm"
)

type LoginHash struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`
	Hash      *string        `gorm:"size:255;" json:"-"`
	UserId    uint           `gorm:"size:255;not null;" json:"user_id"`
	CreatedAt time.Time      `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (loginhash *LoginHash) Save() (*LoginHash, error) {
	err := database.Database.Create(&loginhash).Error
	if err != nil {
		return &LoginHash{}, err
	}
	return loginhash, nil
}

func FindLoginHashByUserId(id uint) (LoginHash, error) {
	var loginhash LoginHash
	err := database.Database.Where("ID=?", id).Find(&loginhash).Error
	if err != nil {
		return LoginHash{}, err
	}
	return loginhash, nil
}
