package gate

import (
	"time"

	"gorm.io/gorm"
)

type AdminScope struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	AdminID   uint           `json:"admin_id" gorm:""`
	ScopeID   uint           `json:"scope_id" gorm:""`
}
