package api

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/service/tasks"
)


type TaskApi struct {}


//
// @Summary 任务列表分页
// @Description
// @Tags 任务
// @Accept   json
// @Produce  json
// @Param body body req.SelectTaskOptionReqWithPage true "The object content"
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /task/page [POST]
func (api *TaskApi) Page(c *gin.Context) {

    var opt *req.SelectTaskOptionReqWithPage
    defer logError(c)
    _ = c.ShouldBindJSON(opt)

    page, data, _ := tasks.Svc().SelectTaskWithPage(c, opt)

    success(c, data, page)

}


//
// @Summary 创建任务
// @Description
// @Tags 任务
// @Accept   json
// @Produce  json
// @Param body body req.TaskDataReq true "The object content"
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /task/create [POST]
func (api *TaskApi) Create(c *gin.Context) {

    var opt *req.TaskDataReq
    defer logError(c)
    _ = c.ShouldBindJSON(opt)

    err := tasks.Svc().CreateTask(c, opt)

    success(c, err)

}



//
// @Summary 更新任务
// @Description
// @Tags 任务
// @Accept   json
// @Produce  json
// @Param body body req.TaskDataReq true "The object content"
// @Success 200 {object}  Response
// @Failure 400 {object}  Response
// @Failure 500 {object}  Response
// @Router /task/update [POST]
func (api *TaskApi) Update(c *gin.Context) {

    var opt *req.TaskDataReq
    defer logError(c)
    _ = c.ShouldBindJSON(opt)

    err := tasks.Svc().UpdateTask(c, opt)

    success(c, err)

}
