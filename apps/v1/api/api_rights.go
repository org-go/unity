package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/service/rights"
)


type RightsApi struct {}


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
// @Router /rights/page [POST]
func (api *RightsApi) Page(c *gin.Context) {

    var opt *req.SelectRightsOptionReqWithPage
    defer logError(c)
    _ = c.ShouldBindJSON(opt)

    data, page, _ := rights.Svc().SelectRightsWithPage(c, opt)

    success(c, data, page)

}
