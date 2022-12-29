package welfare

import (
	"context"
	"unity.service/apps/v1/dto/common"
	"unity.service/apps/v1/dto/req"
	"unity.service/apps/v1/entity/assembler"
)

type (
	//  iWelfareRepoInterface
	//  D_welfare
	//  福利表
	iWelfareRepoInterface interface {

		//  Create
		Create(ctx context.Context, welfare assembler.WelfareAssembler) error

		//  Update
		Update(ctx context.Context, welfare assembler.WelfareAssembler) error

		//	Select
		Select(ctx context.Context, option req.SelectWelfareOptionReqWithPage) (welfares *assembler.WelfareAssemblers, page *common.PaginationResult,  err error)

		//  Query
		Query(ctx context.Context, condition assembler.WelfareAssembler) (welfare *assembler.WelfareAssembler, err error)

	}


	//  iWelfareClassifyRepository
	//  D_welfare_classify
	//  福利类目表
	iWelfareClassifyRepository interface {
		//  Option
		Option(ctx context.Context) (classify *assembler.WelfareClassifyAssemblers)
	}

)
