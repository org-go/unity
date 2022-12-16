package assembler

import (
	"unity.service/apps/v1/dto/req"
	"unity.service/apps/v1/entity/model"
)

type (

	//  WelfareAssemblers
	WelfareAssemblers  []WelfareAssembler
	WelfareAssembler model.DWelfare

)

/**
 * $(OPT)
 * @Description:
 * @receiver o
 * @param data
 * @return WelfareAssembler
 */
func (o *WelfareAssembler) DTO(data *req.WelfareDataReq) WelfareAssembler {
	return *o
}


