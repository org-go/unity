package platforms

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/dto/resp"
)

type (

	//  iPlatformService
	 iPlatformService interface {

		 OptionPlatform(c *gin.Context) resp.OptionsResp
	 }
)
