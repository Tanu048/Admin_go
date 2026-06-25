package models

import "time"

type Admin struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	AdminID      string    `json:"admin_id" gorm:"unique;not null"` 
	CreatedAt    time.Time `json:"created_at"`
	EmailID      string    `json:"email_id" gorm:"unique;not null"`
	PasswordHash string    `json:"-"`
	IsSuperAdmin bool      `json:"is_super_admin" gorm:"default:false"`
}