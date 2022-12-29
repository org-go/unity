package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/service/tasks"
)

type TaskOrderApi struct {}


//
// @Summary 接单(领取任务)
// @Description
// @Tags 任务
// @Accept   json
// @Produce  json
// @Param			task_id	 query		integer		true	"任务单号"
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /task_order/receive [GET]
func (api *TaskOrderApi) Receive(c *gin.Context) {

    var opt *req.TaskOrderReceiveReq
    defer logError(c)
    _ = c.ShouldBindQuery(opt)

    err := tasks.Svc().ReceiveTaskOrder(c, opt)

    success(c, err)

}
