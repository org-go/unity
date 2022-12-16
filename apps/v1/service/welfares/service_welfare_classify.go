package welfares

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/resp"
    "unity.service/apps/v1/repository/welfare"
)

/**
 * $(OptionWelfareClassify)
 * @Description:
 * @receiver this
 * @param c
 * @return options
 */
func (this *welfareService) OptionWelfareClassify(c *gin.Context) (opts resp.OptionsResp) {

    return welfare.WelfareClassify.Option(c).Options()

}

