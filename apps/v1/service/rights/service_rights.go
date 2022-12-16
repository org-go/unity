package rights

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
)

/**
 * $(SelectRightsWithPage)
 * @Description:
 * @receiver this
 * @param c
 * @param opt
 * @return rights
 * @return page
 * @return err
 */
func (this rightsService) SelectRightsWithPage(c *gin.Context, opt *req.SelectRightsOptionReqWithPage) (rights assembler.RightsAssemblers, page *common.PaginationResult, err error) {
    return nil, nil, err
}
