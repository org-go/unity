package req

import "unity.service/apps/v1/dto/common"

//
//  TaskOrderReceiveReq
//  @Description: 领取任务单
//
type TaskOrderReceiveReq struct {
    TaskID int64 `form:"task_id"` // 任务单号
}

//
//  SelectTaskOrderOptionReqWithPage
//  @Description:
//
type SelectTaskOrderOptionReqWithPage struct {
    common.PaginationParam
    SelectTaskOrderOptionReq
}

//  SelectTaskOrderOptionReq
type SelectTaskOrderOptionReq struct {
    TaskID         int64 `json:"task_id"`          // 任务单号
    TaskClassifyPK int   `json:"task_classify_pk"` // 任务类型PK
    Status         byte  `json:"status"`           // 状态：
}
