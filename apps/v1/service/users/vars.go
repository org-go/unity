package users

import (
    "unity.service/apps/v1/entity/assembler"
    "unity.service/apps/v1/service"
)

/**
 * $(Svc)
 * @Description:
 * @return iUserService
 */
func Svc() iUserService {
    return &userService{edge: service.Initialization()}
}

//
//  platformService
//  @Description:
//
type (
    userService struct {
        edge     *service.Edge
        assemble assemble
    }

    assemble struct {
        users *assembler.UserAssemblers
        user  *assembler.UserAssembler
    }
)












