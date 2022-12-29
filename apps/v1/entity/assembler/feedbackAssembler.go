package assembler

import (
	"time"
	"unity.service/apps/v1/dto/req"
	"unity.service/apps/v1/entity/model"
)

type (

	//  FeedbackAssemblers
	FeedbackAssemblers  []FeedbackAssembler
	FeedbackAssembler model.DFeedback

)

/**
 * $(DTO)
 * @Description:
 * @receiver o
 * @param data
 * @return FeedbackAssembler
 */
func (o FeedbackAssembler) DTO(data *req.FeedbackDataReq) FeedbackAssembler {
	return FeedbackAssembler{
		Type:        data.Type,
		Description: data.Description,
		CreatedTime: time.Now(),
		CreatedUser: ``,
	}
}
