package model

import (
	"time"
)

type BaseModel struct {
	ID        int64      `json:"id"`
	UUID      string     `json:"uuid"` //@meta validate_update:true
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
