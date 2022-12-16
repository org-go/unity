package rank

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type rankRepo struct {}


/**
 * $(Option)
 * @Description:
 * @receiver this
 * @param c
 * @return ranks
 */
func (this rankRepo) Option(c *gin.Context) (ranks *assembler.RankAssemblers) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DRank).TableName())
    _ = db.Select(`pk, name`).Find(ranks).Error
    return ranks

}




