package model

import "time"

type DWelfare struct {
	Pk                int64     `gorm:"pk" json:"pk"`
	Id                int64     `gorm:"id" json:"id"`                                   // 对外ID
	PlatformPk        int32     `gorm:"platform_pk" json:"platform_pk"`                 // 平台PK
	WelfareClassifyPk int32     `gorm:"welfare_classify_pk" json:"welfare_classify_pk"` // 福利类型PK
	Name              string    `gorm:"name" json:"name"`                               // 福利名称
	Title             string    `gorm:"title" json:"title"`                             // 福利标题
	Description       string    `gorm:"description" json:"description"`                 // 福利描述
	Link              string    `gorm:"link" json:"link"`                               // 福利链接
	Status            byte      `gorm:"status" json:"status"`                           // 状态; 1:启用； 2：禁用
	CreatedTime       time.Time `gorm:"created_time" json:"created_time"`               // 创建时间
	UpdatedTime       time.Time `gorm:"updated_time" json:"updated_time"`               // 更新时间
	CreatedUser       string    `gorm:"created_user" json:"created_user"`               // 创建人
}

func (*DWelfare) TableName() string {
	return "d_welfare"
}
