package models

import (
	"beew/utils/formater"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt formater.XTime `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt *time.Time     `sql:"index" json:"-"`
}
