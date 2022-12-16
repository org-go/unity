package welfares

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/repository/welfare"
)

/**
 * $(CreateWelfare)
 * @Description:
 * @receiver this
 * @param c
 * @param data
 * @return error
 */
func (this *welfareService) CreateWelfare(c *gin.Context, data *req.WelfareDataReq) error {

    return welfare.Welfare.Create(c, this.assemble.welfare.DTO(data))

}

/**
 * $(UpdateWelfare)
 * @Description:
 * @receiver this
 * @param c
 * @param data
 * @return error
 */
func (this *welfareService) UpdateWelfare(c *gin.Context, data *req.WelfareDataReq) error {

    return welfare.Welfare.Update(c, this.assemble.welfare.DTO(data))

}

/**
 * $(SelectWelfareWithPage)
 * @Description:
 * @receiver this
 * @param c
 * @param option
 * @return page
 * @return welfares
 * @return err
 */
func (this *welfareService) SelectWelfareWithPage(c *gin.Context, option *req.SelectWelfareOptionReqWithPage) (welfares *assembler.WelfareAssemblers, page *common.PaginationResult,  err error) {
    return welfare.Welfare.Select(c, *option)
}
