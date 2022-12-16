package welfares

import (

    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/service"
)

/**
 * $(Svc)
 * @Description:
 * @return iWelfareService
 */
func Svc() iWelfareService {
    return &welfareService{edge: service.Initialization()}
}

//
//  welfareService
//  @Description:
//
type (
    welfareService struct {
        edge     *service.Edge
        assemble assemble
    }

    assemble struct {
        welfares *assembler.WelfareAssemblers
        welfare  *assembler.WelfareAssembler
        welfareClassifys *assembler.WelfareClassifyAssemblers
        welfareClassify *assembler.WelfareClassifyAssembler
    }
)






