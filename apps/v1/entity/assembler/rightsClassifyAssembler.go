package assembler

import (

	"unity.service/apps/v1/dto/resp"
	"unity.service/apps/v1/entity/model"
)

type (

	//  RightsClassifyAssemblers
	RightsClassifyAssemblers  []RightsClassifyAssembler
	RightsClassifyAssembler model.DRightsClassify

)


/**
 * $(Option)
 * @Description:
 * @receiver os
 * @return opts
 */
func (os RightsClassifyAssemblers) Options() (opts resp.OptionsResp) {

	for _, o := range os {
		opts = append(opts, &resp.OptionResp{
			Text:  o.Name,
			Value: o.Pk,
		})
	}
	return opts
}
