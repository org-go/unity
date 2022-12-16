package model

import "time"

type DTaskClassify struct {
	Pk          int32     `gorm:"pk" json:"pk"`
	PlatformPk  int32     `gorm:"platform_pk" json:"platform_pk"`   // 平台PK
	Name        string    `gorm:"name" json:"name"`                 // 任务类目名称
	Description string    `gorm:"description" json:"description"`   // 描述
	Status      byte      `gorm:"status" json:"status"`             // 状态：1：启用， 2：禁用
	CreatedTime time.Time `gorm:"created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"` // 更新时间
	CreatedUser string    `gorm:"created_user" json:"created_user"` // 创建人
}

func (*DTaskClassify) TableName() string {
	return "d_task_classify"
}
