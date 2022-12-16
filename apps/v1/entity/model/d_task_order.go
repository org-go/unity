package model

import "time"

type DTaskOrder struct {
	Pk          int64     `gorm:"pk" json:"pk"`
	WorkerId    int32     `gorm:"worker_id" json:"worker_id"`       // 接单人ID
	TaskId      int64     `gorm:"task_id" json:"task_id"`           // 任务ID
	OrderNumber string    `gorm:"order_number" json:"order_number"` // 任务接单号（任务ID+自定义）
	Score       int32     `gorm:"score" json:"score"`               // 任务积分
	Step        string    `gorm:"step" json:"step"`                 // 完成步骤
	CreatedTime time.Time `gorm:"created_time" json:"created_time"` // 创建接单时间
	UpdatedTime time.Time `gorm:"updated_time" json:"updated_time"` // 更新接单时间
	Evaluate    byte      `gorm:"evaluate" json:"evaluate"`         // 评价 1：未完成 2：未达标 10：已达标
	Auditor     string    `gorm:"auditor" json:"auditor"`           // 评审人
}

func (*DTaskOrder) TableName() string {
	return "d_task_order"
}
