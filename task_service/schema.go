package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type response struct {
	Code    int         `json:"code" example:"200" `
	Message string      `json:"message" example:"Error Message"`
	Data    interface{} `json:"data"`
} //@name Response

// Task Struct for db model
type Task struct {
	// gorm.Model `json:"-"`
	// gorm.Model `gorm:"table:tasks"`
	gorm.Model `swaggerignore:"true"`
	UUID       uuid.UUID `gorm:"type:uuid;primary_key;" swaggerignore:"true"`
	// ID     int    `json:"id" db:"id"`
	Title  string `json:"title" binding:"required,min=3" example:"Title"`
	Status string `json:"status" binding:"required,min=3" example:"todo"`
	UserID int    `json:"user_id" example:"0"`
} //@name Task

// TableName overrides the table name
func (Task) TableName() string {
	return "tasks"
}
