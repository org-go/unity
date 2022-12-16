package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/service/tasks"
)


type TaskApi struct {}


//
// @Summary 任务列表分页
// @Description
// @Tags 任务
// @Accept   json
// @Produce  json
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /task/page [POST]
func (api *TaskApi) Page(c *gin.Context) {

    defer logError(c)
    option := tasks.Svc()

    success(c, option)

}

