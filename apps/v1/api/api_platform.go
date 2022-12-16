package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/service/platforms"
)


type PlatformApi struct {}


//
// @Summary 平台选项
// @Description
// @Tags 平台
// @Accept   json
// @Produce  json
// @Success 200 {object}  Response{data=resp.OptionsResp}
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /platform/option [GET]
func (api *PlatformApi) Option(c *gin.Context) {

    defer logError(c)
    option :=platforms.Svc().OptionPlatform(c)

    success(c, option)

}

