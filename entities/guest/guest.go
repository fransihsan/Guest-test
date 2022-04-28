package guest

import (
	"time"

	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);not null"`
	NoHP  string `gorm:"type:varchar(255);not null;unique"`
	Date  time.Time
	Pesan string `gorm:"type:varchar(255);not null"`
	UserID uint
}
