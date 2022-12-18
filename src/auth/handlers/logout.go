package handlers

import (
	"net/http"
	"time"

	helpers "github.com/MaestroJolly/go-be-api-scaffold/src/db/helpers"
	"github.com/MaestroJolly/go-be-api-scaffold/src/db/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginHash struct {
	ID        uint           `gorm:"primarykey;autoIncrement" json:"id"`
	Hash      string         `gorm:"size:255;" json:"-"`
	UserId    string         `gorm:"size:255;not null;" json:"user_id"`
	CreatedAt time.Time      `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func Logout() gin.HandlerFunc {
	return func(context *gin.Context) {

		_, err := helpers.CurrentUser(context)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		loginhash := models.LoginHash{
			Hash: nil,
		}

		_, err = loginhash.Save()

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "logout successfully",
		})
	}
}
