package task

import (
    "context"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type taskRepo struct {}



/**
 * $(Create)
 * @Description:
 * @receiver this
 * @param ctx
 * @param task
 * @return error
 */
func (this *taskRepo) Create(ctx context.Context, task assembler.TaskAssembler) error {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTask).TableName())
    err := db.Create(task).Error
    return err

}

/**
 * $(Update)
 * @Description:
 * @receiver this
 * @param ctx
 * @param task
 * @return error
 */
func (this *taskRepo) Update(ctx context.Context, task assembler.TaskAssembler) error {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTask).TableName())
    err := db.Save(task).Error
    return err

}


/**
 * $(Select)
 * @Description:
 * @receiver this
 * @param ctx
 * @param option
 * @return tasks
 * @return page
 * @return err
 */
func (this *taskRepo) Select(ctx context.Context, option req.SelectTaskOptionReqWithPage) (tasks *assembler.TaskAssemblers, page *common.PaginationResult, err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTask).TableName())
    if v := option.SelectTaskOptionReq.ID; v > 0 {
        db = db.Where(`id = ?`, v)
    }

    page, err = repo.WrapPageQuery(db, option.PaginationParam, tasks)

    return tasks, page, err
}


/**
 * $(Query)
 * @Description:
 * @receiver this
 * @param ctx
 * @param condition
 * @return task
 * @return err
 */
func (this *taskRepo) Query(ctx context.Context, condition assembler.TaskAssembler) (task *assembler.TaskAssembler, err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTask).TableName())
    if v := condition.Id; v > 0 {
        db = db.Where(`id = ?`, v)
    }
    err = db.Find(task).Error
    return task, err

}
