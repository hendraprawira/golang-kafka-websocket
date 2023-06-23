package models

import "time"

type Client struct {
	ID        uint64    `json:"id" gorm:"primaryKey;auto_increment;not_null"`
	Username  string    `json:"username" binding:"required" `
	Email     string    `json:"email" binding:"required" `
	Fullname  string    `json:"fullname" gorm:"not null" binding:"required"`
	CreatedBy uint64    `json:"created_by" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy uint64    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type ClientModelAdd struct {
	Username string `json:"username" binding:"required" `
	Email    string `json:"email" binding:"required" `
	Fullname string `json:"fullname" gorm:"not null" binding:"required"`
}

func (Client) TableName() string {
	return "clients"
}
