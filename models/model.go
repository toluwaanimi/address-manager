package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
