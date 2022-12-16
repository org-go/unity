package tasks

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/dto/resp"
)

type (

	//  iTaskService
	 iTaskService interface {

		 iTaskClassifyService
	 }

	 //  iTaskClassifyService
	 iTaskClassifyService interface {

		 OptionTaskClassify(c *gin.Context) (opts resp.OptionsResp)

	 }
)
