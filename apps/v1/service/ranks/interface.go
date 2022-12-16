package ranks

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/dto/resp"
)

type (

	//  iRankService
	 iRankService interface {

		 OptionRank(c *gin.Context) (opts resp.OptionsResp)
	 }
)
