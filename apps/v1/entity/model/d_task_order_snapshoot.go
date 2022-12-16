package model

import "time"

type DTaskOrderSnapshoot struct {
	Pk            int64     `gorm:"pk" json:"pk"`
	OrderNumber   string    `gorm:"order_number" json:"order_number"`       // 任务订单号
	TaskOwnerName string    `gorm:"task_owner_name" json:"task_owner_name"` // 任务负责人
	TaskId        int64     `gorm:"task_id" json:"task_id"`                 // 任务ID
	TaskName      string    `gorm:"task_name" json:"task_name"`             // 任务名称
	StartTime     time.Time `gorm:"start_time" json:"start_time"`           // 开始时间
	EndTime       time.Time `gorm:"end_time" json:"end_time"`               // 结束时间
	WorkerId      string    `gorm:"worker_id" json:"worker_id"`             // 接单人
	Claim         string    `gorm:"claim" json:"claim"`                     // 要求
	Step          string    `gorm:"step" json:"step"`                       // 步骤
	Score         int32     `gorm:"score" json:"score"`                     // 积分
	Evaluate      byte      `gorm:"evaluate" json:"evaluate"`               // 评价
	CreatedTime   time.Time `gorm:"created_time" json:"created_time"`       // 创建时间
}

func (*DTaskOrderSnapshoot) TableName() string {
	return "d_task_order_snapshoot"
}
