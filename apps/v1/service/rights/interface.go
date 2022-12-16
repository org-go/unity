package rights

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/dto/common"
	"unity.service/apps/v1/dto/req"
	"unity.service/apps/v1/dto/resp"
	"unity.service/apps/v1/entity/assembler"
)

type (

	//  iRightsService
	 iRightsService interface {

		 SelectRightsWithPage(c *gin.Context, opt *req.SelectRightsOptionReqWithPage) (rights assembler.RightsAssemblers, page *common.PaginationResult, err error)

		iRightsClassifyService
	 }

	 //
	 iRightsClassifyService interface {
		 OptionRightsClassify(c *gin.Context) (opts resp.OptionsResp)
	 }
)
