package tasks

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/errx"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/repository/task"
)

/**
 * $(ReceiveTaskOrder)
 * @Description: 领取任务
 * @receiver this
 * @param c
 * @param opt
 * @return error
 */
func (this taskService) ReceiveTaskOrder(c *gin.Context, opt *req.TaskOrderReceiveReq) error {

    //s1
    info, err := task.Task.Query(c, assembler.TaskAssembler{Id: opt.TaskID})
    if err != nil || info.Pk == 0 {
        return errx.Rquest
    }
    //s2
    err = task.TaskOrder.Create(c, this.assemble.taskOrder.DTO(1, *info))

    return err

}
