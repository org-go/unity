package welfare

import (
    "context"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type welfareRepo struct {

}



/**
 * $(Create)
 * @Description:
 * @receiver this
 * @param ctx
 * @param welfare
 * @return error
 */
func (this *welfareRepo) Create(ctx context.Context, welfare assembler.WelfareAssembler) error {

    db := repo.ServiceDBA.Mysql.Table(new(model.DWelfare).TableName())
    err := db.Create(welfare).Error
    return err

}

/**
 * $(Update)
 * @Description:
 * @receiver this
 * @param ctx
 * @param welfare
 * @return error
 */
func (this *welfareRepo) Update(ctx context.Context, welfare assembler.WelfareAssembler) error {

    db := repo.ServiceDBA.Mysql.Table(new(model.DWelfare).TableName())
    err := db.Save(welfare).Error
    return err

}

/**
 * $(Select)
 * @Description:
 * @receiver this
 * @param ctx
 * @param search
 * @return page
 * @return welfares
 * @return err
 */
func (this *welfareRepo) Select(ctx context.Context, option req.SelectWelfareOptionReqWithPage) (welfares *assembler.WelfareAssemblers, page *common.PaginationResult,  err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DWelfare).TableName())
    if v := option.SelectWelfareOptionReq.ID; v > 0 {
        db = db.Where(`id = ?`, v)
    }
    if v := option.SelectWelfareOptionReq.PlatformID; v > 0 {
        db = db.Where(`platform_pk = ?`, v)
    }
    if v := option.SelectWelfareOptionReq.WelfareClassifyID; v > 0 {
        db = db.Where(`welfare_classify_pk = ?`, v)
    }
    if v := option.SelectWelfareOptionReq.Name; v != `` {
        db = db.Where(`name like ?`, "%"+v+"%")
    }
    if v := option.SelectWelfareOptionReq.Title; v != `` {
        db = db.Where(`title like ?`, "%"+v+"%")
    }
    if v := option.SelectWelfareOptionReq.Status; v  > 0 {
        db = db.Where(`status = ?`, v)
    }
    page, err = repo.WrapPageQuery(db, option.PaginationParam, welfares)

    return welfares, page, err

}


/**
 * $(Query)
 * @Description:
 * @receiver this
 * @param ctx
 * @param condition
 * @return welfare
 * @return err
 */
func (this *welfareRepo) Query(ctx context.Context, condition assembler.WelfareAssembler) (welfare *assembler.WelfareAssembler, err error) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DWelfare).TableName())
    if v := condition.Pk; v > 0 {
        db = db.Where(`pk = ?`, v)
    }
    if v := condition.Id; v > 0 {
        db = db.Where(`id = ?`, v)
    }
    err = db.Find(welfare).Error

    return welfare, err
}
