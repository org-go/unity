package tasks

import (
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/service"
)

/**
 * $(Svc)
 * @Description:
 * @return iWelfareService
 */
func Svc() iTaskService {
    return &taskService{edge: service.Initialization()}
}

//
//  taskService
//  @Description:
//
type (
    taskService struct {
        edge     *service.Edge
        assemble assemble
    }

    assemble struct {
        tasks         *assembler.TaskAssemblers
        task          *assembler.TaskAssembler
        taskClassifys *assembler.TaskClassifyAssemblers
        taskClassify  *assembler.TaskClassifyAssembler
    }
)
