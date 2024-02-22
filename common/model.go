package common

import (
	"time"
)

type Model struct {
	Id        uint      `gorm:"id"`
	AppId     uint      `gorm:"app_id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	CreatedBy string    `gorm:"created_by"`
	UpdatedBy string    `gorm:"updated_by"`
}
