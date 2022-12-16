package task

import (
    "context"
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

