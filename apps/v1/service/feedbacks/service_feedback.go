package feedbacks

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/repository/feedback"
)

/**
 * $(CreateFeedback)
 * @Description: 创建反馈信息
 * @receiver this
 * @param c
 * @param data
 * @return error
 */
func (this *feedbackService) CreateFeedback(c *gin.Context, data *req.FeedbackDataReq) error {

    err := feedback.Feedback.Create(c, this.assemble.feedback.DTO(data))
    return err

}


/**
 * $(SelectFeedbackWithPage)
 * @Description: 查询反馈信息
 * @receiver this
 * @param c
 * @param opt
 * @return fbs
 * @return page
 * @return err
 */
func (this *feedbackService) SelectFeedbackWithPage(c *gin.Context, opt *req.SelectFeedbackOptionReqWithPage) (fbs *assembler.FeedbackAssemblers, page *common.PaginationResult, err error) {

    fbs, page, err = feedback.Feedback.Select(c, *opt)
    return fbs, page, err

}










