package assembler

import (
	"time"
	"unity.service/apps/v1/dto/req"
	"unity.service/apps/v1/entity/model"
	"unity.service/pkg/go/idx"
)

type (

	//  TaskAssemblers
	TaskAssemblers  []TaskAssembler
	TaskAssembler model.DTask

)

/**
 * $(DTO)
 * @Description:
 * @receiver o
 * @param opt
 * @return TaskAssembler
 */
func (o TaskAssembler) DTO(opt *req.TaskDataReq) TaskAssembler {
	return TaskAssembler{
		Id:             idx.Tid(),
		TaskClassifyPk: opt.TaskClassifyPk,
		OwnerUserId:    opt.OwnerUserId,
		Name:           opt.Name,
		Title:          opt.Title,
		Description:    opt.Description,
		Claim:          opt.Claim,
		TaskScore:      opt.TaskScore,
		TaskNumber:     opt.TaskNumber,
		TaskStartTime:  opt.TaskStartTime,
		TaskEndTime:    opt.TaskEndTime,
		Status:         1,
		CreatedTime:    time.Now(),
		UpdatedTime:    time.Now(),
	}
}
