package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/service/welfares"
)


type WelfareClassifyApi struct {}


//
// @Summary 福利类型选项
// @Description
// @Tags 福利
// @Accept   json
// @Produce  json
// @Success 200 {object}  Response{data=resp.OptionsResp}
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /welfare_classify/option [GET]
func (api *WelfareClassifyApi) Option(c *gin.Context) {

    defer logError(c)
    option := welfares.Svc().OptionWelfareClassify(c)

    success(c, option)

}
