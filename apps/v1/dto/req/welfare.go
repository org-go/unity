package req

import (
    "unity.service/apps/v1/dto/common"
)

//
//  SelectWelfareOptionReqWithPage
//  @Description:
//
type SelectWelfareOptionReqWithPage struct {
    common.PaginationParam
    SelectWelfareOptionReq
}
//  SelectWelfareOptionReq
type SelectWelfareOptionReq struct {
    ID                int64  `json:"id"`
    PlatformID        int    `json:"platform_id"`         // 平台ID
    WelfareClassifyID int    `json:"welfare_classify_id"` // 福利类型PK
    Name              string `json:"name"`                // 福利名称
    Title             string `json:"title"`               // 福利标题
    Status            byte   `json:"status"`              // 状态：1:启用； 2：禁用
}


type WelfareDataReq struct{}
