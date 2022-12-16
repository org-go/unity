package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/service/users"
)


type UserApi struct {}


//
// @Summary 用户列表分页
// @Description
// @Tags 用户
// @Accept   json
// @Produce  json
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /user/page [POST]
func (api *UserApi) Page(c *gin.Context) {

    defer logError(c)
    option := users.Svc()

    success(c, option)

}

