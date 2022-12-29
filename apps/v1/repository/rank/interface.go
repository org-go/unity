package rank

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/entity/assembler"
)

type (
	//  iRankRepoInterface
	//  D_rank
	//  头衔表
	iRankRepoInterface interface {
		//  Option
		Option(c *gin.Context) (ranks *assembler.RankAssemblers)
	}


)
