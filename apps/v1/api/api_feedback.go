package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/service/feedbacks"
)


type FeedbackApi struct {}


//
// @Summary 反馈信息列表分页
// @Description
// @Tags 反馈&帮助
// @Accept   json
// @Produce  json
// @Param body body req.SelectFeedbackOptionReqWithPage true "The object content"
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /feedback/page [POST]
func (api *FeedbackApi) Page(c *gin.Context) {

    var opt *req.SelectFeedbackOptionReqWithPage
    defer logError(c)
    _ = c.ShouldBindJSON(opt)

    page, data, _ := feedbacks.Svc().SelectFeedbackWithPage(c, opt)

    success(c, data, page)

}


//
// @Summary 创建反馈信息
// @Description
// @Tags 反馈&帮助
// @Accept   json
// @Produce  json
// @Param body body req.FeedbackDataReq true "The object content"
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /feedback/create [POST]
func (api *FeedbackApi) Create(c *gin.Context) {

    var opt *req.FeedbackDataReq
    defer logError(c)
    _ = c.ShouldBindJSON(opt)

    err := feedbacks.Svc().CreateFeedback(c, opt)

    success(c, err)

}

