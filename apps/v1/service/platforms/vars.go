package platforms

import (
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/service"
)

/**
 * $(Svc)
 * @Description:
 * @return iPlatformService
 */
func Svc() iPlatformService {
    return &platformService{edge: service.Initialization()}
}

//
//  platformService
//  @Description:
//
type (
    platformService struct {
        edge     *service.Edge
        assemble assemble
    }

    assemble struct {
        platforms *assembler.PlatformAssemblers
        platform  *assembler.PlatformAssembler
    }
)









