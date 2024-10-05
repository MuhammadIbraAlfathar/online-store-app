package schema

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id          int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(230);not null"`
	Description string         `json:"description" gorm:"type:text"`
	Price       int            `json:"price" gorm:"not null"`
	Stock       int            `json:"stock" gorm:"not null"`
	CategoryId  int            `json:"category_id" gorm:"not null"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
