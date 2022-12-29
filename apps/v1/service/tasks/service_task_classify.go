package tasks

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/resp"
    "unity.service/apps/v1/repository/task"
)

/**
 * $(OptionTaskClassify)
 * @Description:
 * @receiver t
 * @param c
 * @return opts
 */
func (this taskService) OptionTaskClassify(c *gin.Context) (opts resp.OptionsResp) {

    return task.TaskClassify.Option(c).Options()

}
