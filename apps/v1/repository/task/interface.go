package task

import (
	"context"
	"unity.service/apps/v1/entity/assembler"
)

type (
	//  iTaskRepoInterface
	iTaskRepoInterface interface {

		//  Create
		Create(ctx context.Context, task assembler.TaskAssembler) error

		//  Update
		Update(ctx context.Context, task assembler.TaskAssembler) error


	}


	//  iTaskClassifyRepository
	iTaskClassifyRepository interface {
		//  Option
		Option(ctx context.Context) (classify *assembler.TaskClassifyAssemblers)
	}

)
