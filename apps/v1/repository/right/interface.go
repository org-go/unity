package right

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/entity/assembler"
)

type (
	//  iRightsClassifyRepoInterface
	//  D_rights_classify
	//  权益类目表
	iRightsClassifyRepoInterface interface {
		//  Option
		Option(c *gin.Context) (ranks *assembler.RankAssemblers)
	}


)
