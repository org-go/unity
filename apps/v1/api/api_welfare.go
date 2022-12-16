package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/service/welfares"
)


type WelfareApi struct {}


//
// @Summary 福利列表分页
// @Description
// @Tags 福利
// @Accept   json
// @Produce  json
// @Param body body req.SelectWelfareOptionReqWithPage true "The object content"
// @Success 200 {object}  Response{data=assembler.WelfareAssemblers}
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /welfare/page [POST]
func (api *WelfareApi) Page(c *gin.Context) {

    var opt *req.SelectWelfareOptionReqWithPage
    defer logError(c)
    _ = c.ShouldBindJSON(opt)

    page, data, _ := welfares.Svc().SelectWelfareWithPage(c, opt)

    success(c, data, page)

}
