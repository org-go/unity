package assembler

import (
	"time"
	"unity.service/apps/v1/entity/model"
	"unity.service/pkg/go/idx"
)

type (

	//  TaskOrderAssemblers
	TaskOrderAssemblers  []TaskOrderAssembler
	TaskOrderAssembler model.DTaskOrder

)

/**
 * $(DTO)
 * @Description:
 * @receiver o
 * @param uID
 * @param opt
 * @return TaskOrderAssembler
 */
func (o TaskOrderAssembler) DTO(uID int32, opt TaskAssembler) TaskOrderAssembler {
	return TaskOrderAssembler{
		WorkerId:    uID,
		TaskId:      opt.Id,
		OrderNumber: idx.Sid(),
		Score:       opt.TaskScore,
		CreatedTime: time.Now(),
		Evaluate:    1,
	}
}
