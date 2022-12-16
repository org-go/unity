package model

import "time"

type DRights struct {
	Pk               int32     `gorm:"pk" json:"pk"`
	Name             string    `gorm:"name" json:"name"`                             // 权益名称
	RightsClassifyPk int32     `gorm:"rights_classify_pk" json:"rights_classify_pk"` // 权益类目PK
	Description      string    `gorm:"description" json:"description"`               // 权益描述
	Status           byte      `gorm:"status" json:"status"`                         // 权益状态 1:启用 2：禁用
	CreatedTime      time.Time `gorm:"created_time" json:"created_time"`             // 创建时间
	UpdatedTime      time.Time `gorm:"updated_time" json:"updated_time"`             // 更新时间
}

func (*DRights) TableName() string {
	return "d_rights"
}
