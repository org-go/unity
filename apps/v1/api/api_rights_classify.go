package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/service/ranks"
)


type RightsClassifyApi struct {}


//
// @Summary 权益类型选项
// @Description
// @Tags 权益
// @Accept   json
// @Produce  json
// @Success 200 {object}  Response{data=resp.OptionsResp}
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /rights_classify/option [GET]
func (api *RightsClassifyApi) Option(c *gin.Context) {

    defer logError(c)
    option := ranks.Svc().OptionRank(c)

    success(c, option)

}

