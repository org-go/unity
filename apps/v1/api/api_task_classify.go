package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/service/tasks"
)


type TaskClassifyApi struct {}


//
// @Summary 任务类型选项
// @Description
// @Tags 任务
// @Accept   json
// @Produce  json
// @Success 200 {object}  Response{data=resp.OptionsResp}
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /task_classify/option [GET]
func (api *TaskClassifyApi) Option(c *gin.Context) {

    defer logError(c)
    option := tasks.Svc().OptionTaskClassify(c)

    success(c, option)

}
