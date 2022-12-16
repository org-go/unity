package model

import "time"

type DRank struct {
	Pk          int32     `gorm:"pk" json:"pk"`
	Name        string    `gorm:"name" json:"name"`                 // 头衔名称
	Experience  int64     `gorm:"experience" json:"experience"`     // 晋级所需经验
	Description string    `gorm:"description" json:"description"`   // 头衔描述
	RightsPks   string    `gorm:"rights_pks" json:"rights_pks"`     // 权益PK, 逗号(,)分隔
	CreatedTime time.Time `gorm:"created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"` // 更新时间
}

func (*DRank) TableName() string {
	return "d_rank"
}
