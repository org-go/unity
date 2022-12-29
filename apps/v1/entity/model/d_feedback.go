package model

import "time"

type DFeedback struct {
	Pk          int64     `gorm:"pk" json:"pk"`
	Type        byte      `gorm:"type" json:"type"` // 1:意见; 2:建议
	Description string    `gorm:"description" json:"description"`
	CreatedTime time.Time `gorm:"created_time" json:"created_time"`
	CreatedUser string    `gorm:"created_user" json:"created_user"`
}

func (*DFeedback) TableName() string {
	return "d_feedback"
}
