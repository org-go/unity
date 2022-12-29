package task

import (
    "context"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
)

type (
    //  iTaskRepoInterface
    //  D_task
    //  任务表
    iTaskRepoInterface interface {

        //  Create
        Create(ctx context.Context, task assembler.TaskAssembler) error

        //  Update
        Update(ctx context.Context, task assembler.TaskAssembler) error

        //	Select
        Select(ctx context.Context, option req.SelectTaskOptionReqWithPage) (tasks *assembler.TaskAssemblers, page *common.PaginationResult, err error)

        //  Query
        Query(ctx context.Context, condition assembler.TaskAssembler) (task *assembler.TaskAssembler, err error)

    }

    //  iTaskOrderRepoInterface
    //  D_task_order
    //  任务订单表
    iTaskOrderRepoInterface interface {

        //  Create
        Create(ctx context.Context, order assembler.TaskOrderAssembler) error

        //  Update
        Update(ctx context.Context, order assembler.TaskOrderAssembler) error

        //	Select
        Select(ctx context.Context, option req.SelectTaskOrderOptionReqWithPage) (orders *assembler.TaskOrderAssemblers, page *common.PaginationResult, err error)

        //  Query
        Query(ctx context.Context, condition assembler.TaskOrderAssembler) (order *assembler.TaskOrderAssembler, err error)

    }


    //  iTaskClassifyRepository
    //  D_task_classify
    //  任务类目表
    iTaskClassifyRepository interface {
        //  Option
        Option(ctx context.Context) (classify *assembler.TaskClassifyAssemblers)
    }
)
