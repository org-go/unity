package feedbacks

import (
	"github.com/gin-gonic/gin"
	"unity.service/apps/v1/dto/common"
	"unity.service/apps/v1/dto/req"
	"unity.service/apps/v1/entity/assembler"
)

type (

	//  iFeedbackService
	//  feedback
	//  反馈
	iFeedbackService interface {

		 //  CreateFeedback
		 CreateFeedback(c *gin.Context, data *req.FeedbackDataReq) error

		 //  SelectFeedbackWithPage
		 SelectFeedbackWithPage(c *gin.Context, opt *req.SelectFeedbackOptionReqWithPage) (fbs *assembler.FeedbackAssemblers, page *common.PaginationResult, err error)


	}
)
