package feedback

import (
	"context"
	"unity.service/apps/v1/dto/common"
	"unity.service/apps/v1/dto/req"
	"unity.service/apps/v1/entity/assembler"
)

type (
	//  iFeedbackRepoInterface
	//  D_feedback
	//  反馈表
	iFeedbackRepoInterface interface {
		//  Create
		Create(ctx context.Context, feedback assembler.FeedbackAssembler) error

		//	Select
		Select(ctx context.Context, option req.SelectFeedbackOptionReqWithPage) (fbs *assembler.FeedbackAssemblers, page *common.PaginationResult,  err error)

		//  Query
		Query(ctx context.Context, condition assembler.FeedbackAssembler) (fb *assembler.FeedbackAssembler, err error)
	}


)
