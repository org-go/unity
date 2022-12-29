package req

import (
    "time"
    "unity.service/apps/v1/dto/common"
)

//
//  SelectTaskOptionReqWithPage
//  @Description:
//
type SelectTaskOptionReqWithPage struct {
    common.PaginationParam
    SelectTaskOptionReq
}

//  SelectTaskOptionReq
type SelectTaskOptionReq struct {
    ID             int64 `json:"id"`
    TaskClassifyPK int   `json:"task_classify_pk"` // 任务类型PK
    Status         byte  `json:"status"`           // 状态：
}


//
//  TaskDataReq
//  @Description: 创建、更新数据源
//
type TaskDataReq struct{

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
}
