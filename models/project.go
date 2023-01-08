package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Company     string         `gorm:"not null; default:null" json:"company"`
	About       string         `gorm:"not null; type:text; default:null" json:"about"`
	Link        *string        `json:"link"`
	IsPublished bool           `gorm:"default:null; not null" json:"is_published"`
}
