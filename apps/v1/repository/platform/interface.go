package platform

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/entity/assembler"
)

type (
	//  iPlatformRepoInterface
	//  D_platform
	//  平台表
	iPlatformRepoInterface interface {
		//  Option
		Option(c *gin.Context) (platform *assembler.PlatformAssemblers)
	}


)
