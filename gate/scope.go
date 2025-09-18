package gate

import (
	"time"

	"gorm.io/gorm"
)

type Scope struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title       string         `json:"title" gorm:"size:64"`
	Description string         `json:"description" gorm:"size:128"`
}
