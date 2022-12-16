package req

import (
    "unity.service/apps/v1/dto/common"
)

//
//  SelectRightsOptionReqWithPage
//  @Description:
//
type SelectRightsOptionReqWithPage struct {
    common.PaginationParam
    SelectRightsOptionReq
}

//  SelectRightsOptionReq
type SelectRightsOptionReq struct {
    PK               int64 `json:"pk"`
    RightsClassifyID int   `json:"rights_classify_id"` // 福利类型PK
    Status           byte  `json:"status"`             // 状态：1:启用； 2：禁用
}
