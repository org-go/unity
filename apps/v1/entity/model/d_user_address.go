package model

import "time"

type DUserAddress struct {
	Pk          int64     `gorm:"pk" json:"pk"`
	UserPk      int64     `gorm:"user_pk" json:"user_pk"`           // 用户PK
	StateCode   int32     `gorm:"state_code" json:"state_code"`     // 省
	CityCode    int32     `gorm:"city_code" json:"city_code"`       // 市
	StreetCode  int32     `gorm:"street_code" json:"street_code"`   // 街道
	Address     string    `gorm:"address" json:"address"`           // 详细地址
	PostalCode  int32     `gorm:"postal_code" json:"postal_code"`   // 邮编
	CreatedTime time.Time `gorm:"created_time" json:"created_time"` // 创建时间
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"` // 更新时间
	IsDefault   byte      `gorm:"is_default" json:"is_default"`     // 1: 默认
}

func (*DUserAddress) TableName() string {
	return "d_user_address"
}
