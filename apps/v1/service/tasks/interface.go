package tasks

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/dto/resp"
    "unity.service/apps/v1/entity/assembler"
)

type (

    //  iTaskService
    //  Task
    //  任务
    iTaskService interface {
        //  CreateTask
        CreateTask(c *gin.Context, data *req.TaskDataReq) error
        //  UpdateTask
        UpdateTask(c *gin.Context, data *req.TaskDataReq) error
        //  SelectTaskWithPage
        SelectTaskWithPage(c *gin.Context, opt *req.SelectTaskOptionReqWithPage) (tasks *assembler.TaskAssemblers, page *common.PaginationResult, err error)
        //  iTaskOrderService
        iTaskOrderService
        //  iTaskClassifyService
        iTaskClassifyService
    }

    //  iTaskOrderService
    //  Task Order
    //  领取任务
    iTaskOrderService interface {
        ReceiveTaskOrder(c *gin.Context, opt *req.TaskOrderReceiveReq) error
    }

    //  iTaskClassifyService
    //  Task Classify
    //  任务类型
    iTaskClassifyService interface {
        OptionTaskClassify(c *gin.Context) (opts resp.OptionsResp)
    }
)
