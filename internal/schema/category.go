package schema

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"type:varchar(230);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Products  []Product      `json:"products" gorm:"foreignKey:CategoryId"`
}
