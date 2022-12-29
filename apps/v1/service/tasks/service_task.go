package tasks

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/repository/task"
)

/**
 * $(CreateTask)
 * @Description: 创建任务单
 * @receiver this
 * @param c
 * @param task
 * @return error
 */
func (this taskService) CreateTask(c *gin.Context, data *req.TaskDataReq) error {

    return task.Task.Create(c, this.assemble.task.DTO(data))

}

/**
 * $(UpdateTask)
 * @Description: 更新任务单
 * @receiver this
 * @param c
 * @param task
 * @return error
 */
func (this taskService) UpdateTask(c *gin.Context, data *req.TaskDataReq) error {

    return task.Task.Update(c, this.assemble.task.DTO(data))

}

/**
 * $(SelectTaskWithPage)
 * @Description: 任务查询
 * @receiver this
 * @param c
 * @param opt
 * @return tasks
 * @return page
 * @return err
 */
func (this taskService) SelectTaskWithPage(c *gin.Context, opt *req.SelectTaskOptionReqWithPage) (tasks *assembler.TaskAssemblers, page *common.PaginationResult, err error) {

    return task.Task.Select(c, *opt)

}
