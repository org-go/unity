package welfares

import (
    "github.com/gin-gonic/gin"
    "unity.service/apps/v1/dto/common"
    "unity.service/apps/v1/dto/req"
    "unity.service/apps/v1/dto/resp"
    "unity.service/apps/v1/entity/assembler"
)

type (

    //  iWelfareService
    //  welfare
    //  福利
    iWelfareService interface {

        //  CreateWelfare
        CreateWelfare(c *gin.Context, data *req.WelfareDataReq) error
        //  UpdateWelfare
        UpdateWelfare(c *gin.Context, data *req.WelfareDataReq) error
        //  SelectWelfareWithPage
        SelectWelfareWithPage(c *gin.Context, opt *req.SelectWelfareOptionReqWithPage) (welfares *assembler.WelfareAssemblers, page *common.PaginationResult, err error)

        //  iWelfareClassifyService
        iWelfareClassifyService
    }

    //  Classify
    //  welfare classify
    //  福利类目
    iWelfareClassifyService interface {
        OptionWelfareClassify(c *gin.Context) (opts resp.OptionsResp)
    }
)
