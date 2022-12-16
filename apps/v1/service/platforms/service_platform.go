package platforms

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/resp"
    "unity.service/apps/v1/repository/platform"
)


/**
 * $(OptionPlatform)
 * @Description:
 * @receiver this
 * @param c
 * @return resp.OptionsResp
 */
func (this platformService) OptionPlatform(c *gin.Context) (opts resp.OptionsResp) {

    return platform.Platform.Option(c).Options()

}
