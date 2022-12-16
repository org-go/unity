package coreSystem

import (
    "unity.service/apps/v1/entity"
)

type (
    core struct {
        config *CoreConfig
        optionCore
        help
    }
    CoreConfig struct {
        Appid     string `json:"appid"`
        Appsecret string `json:"appsecret"`
        Site      string `json:"site"`
        Url       string `json:"url"`
    }

    optionCore struct {
        Authorized entity.Option
        Payment    entity.Option
    }

    help struct{}
)

func RegistryCorePlatform(config *CoreConfig) IcorePlatform {
    return &core{config: config}
}

func (this core) option() *optionCore {
    return &optionCore{

        // 支付/退款
        Payment: entity.Option{
            Request: "POST",
            Uri:     this.config.Site + this.config.Url,
        },
    }
}
