package ranks

import (
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/service"
)

/**
 * $(Svc)
 * @Description:
 * @return iRankService
 */
func Svc() iRankService {
    return &rankService{edge: service.Initialization()}
}

//
//  platformService
//  @Description:
//
type (
    rankService struct {
        edge     *service.Edge
        assemble assemble
    }

    assemble struct {
        ranks *assembler.RankAssemblers
        rank  *assembler.RankAssembler
    }
)












