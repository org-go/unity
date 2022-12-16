package right

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type rightsClassifyRepo struct {}


/**
 * $(Option)
 * @Description:
 * @receiver this
 * @param c
 * @return ranks
 */
func (this rightsClassifyRepo) Option(c *gin.Context) (ranks *assembler.RankAssemblers) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DRightsClassify).TableName())
    _ = db.Select(`pk, name`).Find(ranks).Error
    return ranks

}




