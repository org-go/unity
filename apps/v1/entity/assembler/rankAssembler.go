package assembler

import (

	"unity.service/apps/v1/dto/resp"
	"unity.service/apps/v1/entity/model"
)

type (

	//  RankAssemblers
	RankAssemblers  []RankAssembler
	RankAssembler model.DRank

)


/**
 * $(Option)
 * @Description:
 * @receiver os
 * @return opts
 */
func (os RankAssemblers) Options() (opts resp.OptionsResp) {

	for _, o := range os {
		opts = append(opts, &resp.OptionResp{
			Text:  o.Name,
			Value: o.Pk,
		})
	}
	return opts
}
