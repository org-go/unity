package model

import "time"

type DTask struct {
	Pk             int64     `gorm:"pk" json:"pk"`
	Id             int64     `gorm:"id" json:"id"`                             // 任务ID
	TaskClassifyPk int32     `gorm:"task_classify_pk" json:"task_classify_pk"` // 任务类目PK
	OwnerUserId    int64     `gorm:"owner_user_id" json:"owner_user_id"`       // 任务负责人
	Name           string    `gorm:"name" json:"name"`                         // 名称
	Title          string    `gorm:"title" json:"title"`                       // 标题
	Description    string    `gorm:"description" json:"description"`           // 描述
	Claim          string    `gorm:"claim" json:"claim"`                       // 要求
	TaskScore      int32     `gorm:"task_score" json:"task_score"`             // 任务积分/单
	TaskNumber     int32     `gorm:"task_number" json:"task_number"`           // 单数
	TaskStartTime  time.Time `gorm:"task_start_time" json:"task_start_time"`   // 任务开始时间
	TaskEndTime    time.Time `gorm:"task_end_time" json:"task_end_time"`       // 任务结束时间
	Status         byte      `gorm:"status" json:"status"`                     // 状态：1:草稿， 3：审核中， 5：进行中， 7：下架， 9：结束
	CreatedTime    time.Time `gorm:"created_time" json:"created_time"`         // 创建时间
	UpdatedTime    time.Time `gorm:"updated_time" json:"updated_time"`         // 更新时间
	Auditor        string    `gorm:"auditor" json:"auditor"`                   // 审核人
}

func (*DTask) TableName() string {
	return "d_task"
}
