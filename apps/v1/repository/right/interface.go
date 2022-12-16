package right

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/entity/assembler"
)

type (
	//  iRightsClassifyRepoInterface
	iRightsClassifyRepoInterface interface {
		//  Option
		Option(c *gin.Context) (ranks *assembler.RankAssemblers)
	}


)
