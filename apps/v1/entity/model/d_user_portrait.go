package model

type DUserPortrait struct {
	UserPk  int64  `gorm:"user_pk" json:"user_pk"`
	UnitId  string `gorm:"unit_id" json:"unit_id"` // 身份证
	Borth   string `gorm:"borth" json:"borth"`     // 生日
	Hobbies string `gorm:"hobbies" json:"hobbies"` // 兴趣爱好
	Email   string `gorm:"email" json:"email"`     // 邮箱
}

func (*DUserPortrait) TableName() string {
	return "d_user_portrait"
}
