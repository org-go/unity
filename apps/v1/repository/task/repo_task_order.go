package task

import (
    "context"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type taskOrderRepo struct {}



/**
 * $(Create)
 * @Description:
 * @receiver this
 * @param ctx
 * @param order
 * @return error
 */
func (this *taskOrderRepo) Create(ctx context.Context, order assembler.TaskOrderAssembler) error {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTaskOrder).TableName())
    err := db.Create(order).Error
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
func (this *taskOrderRepo) Update(ctx context.Context, task assembler.TaskOrderAssembler) error {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTaskOrder).TableName())
    err := db.Save(task).Error
    return err

}


/**
 * $(Select)
 * @Description:
 * @receiver this
 * @param ctx
 * @param option
 * @return orders
 * @return page
 * @return err
 */
func (this *taskOrderRepo) Select(ctx context.Context, option req.SelectTaskOrderOptionReqWithPage) (orders *assembler.TaskOrderAssemblers, page *common.PaginationResult, err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTaskOrder).TableName())
    if v := option.SelectTaskOrderOptionReq.TaskID; v > 0 {
        db = db.Where(`task_id = ?`, v)
    }

    page, err = repo.WrapPageQuery(db, option.PaginationParam, orders)

    return orders, page, err
}


/**
 * $(Query)
 * @Description:
 * @receiver this
 * @param ctx
 * @param condition
 * @return order
 * @return err
 */
func (this *taskOrderRepo) Query(ctx context.Context, condition assembler.TaskOrderAssembler) (order *assembler.TaskOrderAssembler, err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTask).TableName())
    if v := condition.Pk; v > 0 {
        db = db.Where(`pk = ?`, v)
    }
    err = db.Find(order).Error
    return order, err

}
