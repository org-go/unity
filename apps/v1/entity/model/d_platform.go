package model

import "time"

type DPlatform struct {
	Pk          int32     `gorm:"pk" json:"pk"`
	Name        string    `gorm:"name" json:"name"`
	Status      byte      `gorm:"status" json:"status"`
	CreatedTime time.Time `gorm:"created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"`
}

func (*DPlatform) TableName() string {
	return "d_platform"
}
