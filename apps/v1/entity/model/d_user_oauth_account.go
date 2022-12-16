package model

type DUserOauthAccount struct {
	UserPk   int64  `gorm:"user_pk" json:"user_pk"`
	Wechat   string `gorm:"wechat" json:"wechat"`
	Qq       string `gorm:"qq" json:"qq"`
	Dingtalk string `gorm:"dingtalk" json:"dingtalk"`
	Github   string `gorm:"github" json:"github"`
}

func (*DUserOauthAccount) TableName() string {
	return "d_user_oauth_account"
}
