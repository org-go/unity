package req

import (
    "unity.service/apps/v1/dto/common"
)

//
//  SelectFeedbackOptionReqWithPage
//  @Description:
//
type SelectFeedbackOptionReqWithPage struct {
    common.PaginationParam
    SelectFeedbackOptionReq
}

//  SelectFeedbackOptionReq
type SelectFeedbackOptionReq struct {
    Type byte `json:"type"` // 1:意见; 2:建议
}

//
//  FeedbackDataReq
//  @Description: 创建信息
//
type FeedbackDataReq struct {
    Type        byte   `json:"type"`        // 1:意见; 2:建议
    Description string `json:"description"` // 描述
}
