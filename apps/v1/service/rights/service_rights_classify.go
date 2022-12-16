package rights

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/resp"
    "unity.service/apps/v1/repository/right"
)

/**
 * $(OptionRightsClassify)
 * @Description:
 * @receiver this
 * @param c
 * @return opts
 */
func (this rightsService) OptionRightsClassify(c *gin.Context) (opts resp.OptionsResp) {
    return right.RightClassify.Option(c).Options()
}
