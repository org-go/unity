package task

import (
    "context"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type taskClassifyRepo struct {}

/**
 * $(Option)
 * @Description:
 * @receiver this
 * @param ctx
 * @return classify
 */
func (this taskClassifyRepo) Option(ctx context.Context) (classify *assembler.TaskClassifyAssemblers) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DTaskClassify).TableName())
    _ = db.Where(`status = ?`, 1 ).Select(`pk, name`).Find(classify).Error
    return classify

}




