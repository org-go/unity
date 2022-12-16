package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/service/ranks"
)


type RankApi struct {}


//
// @Summary 头衔选项
// @Description
// @Tags 头衔
// @Accept   json
// @Produce  json
// @Success 200 {object}  Response{data=resp.OptionsResp}
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /rank/option [GET]
func (api *RankApi) Option(c *gin.Context) {

    defer logError(c)
    option := ranks.Svc().OptionRank(c)

    success(c, option)

}

