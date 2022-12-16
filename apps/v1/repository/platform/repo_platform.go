package platform

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/entity/model"
    "unity.service/apps/v1/repository/repo"
)

type platformRepo struct {}


/**
 * $(Option)
 * @Description:
 * @receiver this
 * @param c
 * @return platform
 */
func (this platformRepo) Option(c *gin.Context) (platform *assembler.PlatformAssemblers) {

    db := repo.ServiceDBA.Mysql.Table(new(model.DPlatform).TableName())
    _ = db.Where(`status = ?`, 1 ).Select(`pk, name`).Find(platform).Error
    return platform

}




