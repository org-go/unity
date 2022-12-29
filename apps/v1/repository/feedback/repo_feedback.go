package feedback

import (
    "context"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type feedbackRepo struct {}


/**
 * $(Create)
 * @Description:
 * @receiver this
 * @param ctx
 * @param feedback
 * @return error
 */
func (this feedbackRepo) Create(ctx context.Context, feedback assembler.FeedbackAssembler) error {

    db := repo.ServiceDBA.Mysql.Table(new(model.DFeedback).TableName())
    err := db.Save(feedback).Error
    return err

}

/**
 * $(Select)
 * @Description:
 * @receiver this
 * @param ctx
 * @param option
 * @return fbs
 * @return page
 * @return err
 */
func (this feedbackRepo) Select(ctx context.Context, option req.SelectFeedbackOptionReqWithPage) (fbs *assembler.FeedbackAssemblers, page *common.PaginationResult, err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DFeedback).TableName())
    if v := option.Type; v > 0 {
        db = db.Where(`type = ?`, v)
    }
    page, err = repo.WrapPageQuery(db, option.PaginationParam, fbs)

    return fbs, page, err

}

/**
 * $(Query)
 * @Description:
 * @receiver this
 * @param ctx
 * @param condition
 * @return fbs
 * @return err
 */
func (this feedbackRepo) Query(ctx context.Context, condition assembler.FeedbackAssembler) (fb *assembler.FeedbackAssembler, err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DFeedback).TableName())
    if v := condition.Pk; v > 0 {
        db = db.Where(`pk = ?`, v)
    }
    err = db.Find(fb).Error
    return fb, err

}



