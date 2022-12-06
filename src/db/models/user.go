package models

type User struct {
	ID       string `json:"id" binding:"required"`
	UserName string `json:"userName"`
	Email    string `json:"email" binding:"required"`
}
