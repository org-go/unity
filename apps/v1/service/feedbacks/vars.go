package feedbacks

import (
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/service"
)

/**
 * $(Svc)
 * @Description:
 * @return iPlatformService
 */
func Svc() iFeedbackService {
    return &feedbackService{edge: service.Initialization()}
}

//
//  feedbackService
//  @Description:
//
type (
    feedbackService struct {
        edge     *service.Edge
        assemble assemble
    }

    assemble struct {
        feedbacks *assembler.FeedbackAssemblers
        feedback  *assembler.FeedbackAssembler
    }
)

