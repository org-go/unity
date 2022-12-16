package rights

import (
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/service"
)

/**
 * $(Svc)
 * @Description:
 * @return iRankService
 */
func Svc() iRightsService {
    return &rightsService{edge: service.Initialization()}
}

//
//  platformService
//  @Description:
//
type (
    rightsService struct {
        edge     *service.Edge
        assemble assemble
    }

    assemble struct {
        rights *assembler.RightsAssemblers
        right  *assembler.RightsAssembler
        rightsClassifys *assembler.RightsClassifyAssemblers
        rightsClassify *assembler.RightsClassifyAssembler
    }
)


















