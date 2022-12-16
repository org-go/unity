package task

import (
    "context"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type welfareClassifyRepo struct {}

/**
 * $(Option)
 * @Description:
 * @receiver this
 * @param ctx
 * @return classify
 */
func (this welfareClassifyRepo) Option(ctx context.Context) (classify *assembler.WelfareClassifyAssemblers) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DWelfareClassify).TableName())
    _ = db.Where(`status = ?`, 1 ).Select(`pk, name`).Find(classify).Error
    return classify

}




