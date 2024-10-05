package schema

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string         `json:"email" gorm:"type:varchar(230);unique;not null"`
	Name      string         `json:"name" gorm:"type:varchar(230);not null"`
	Address   string         `json:"address" gorm:"type:text"`
	Password  string         `gorm:"type:varchar(230);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
