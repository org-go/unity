package model

import "time"

type DUser struct {
    Pk            int64     `gorm:"pk" json:"pk"`
    Id            int64     `gorm:"id" json:"id"`                           // 对外ID
    SupperiorPk   int64     `gorm:"supperior_pk" json:"supperior_pk"`       // 上级PK
    PathPk        string    `gorm:"path_pk" json:"path_pk"`                 // 级联PK
    Phone         int64     `gorm:"phone" json:"phone"`                     // 手机号
    Account       string    `gorm:"account" json:"account"`                 // 账号
    Password      string    `gorm:"password" json:"password"`               // 密码
    Name          string    `gorm:"name" json:"name"`                       // 名称
    Rank          int32     `gorm:"rank" json:"rank"`                       // 权益级别
    Score         int64     `gorm:"score" json:"score"`                     // 当前积分
    Scope         string    `gorm:"scope" json:"scope"`                     // x： 正常， xx： 禁用，x21： 交易冻结， x22： 任务冻结
    CreatedTime   time.Time `gorm:"created_time" json:"created_time"`       // 创建时间
    UpdatedTime   time.Time `gorm:"updated_time" json:"updated_time"`       // 更新时间
    LastLoginTime time.Time `gorm:"last_login_time" json:"last_login_time"` // 最后登陆时间
    LastLoginIp   string    `gorm:"last_login_ip" json:"last_login_ip"`     // 最后登陆IP
}

func (*DUser) TableName() string {
    return "d_user"
}
