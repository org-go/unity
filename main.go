package main

import (
	_ "go.uber.org/automaxprocs"
	"runtime"
	"unity.service/cli"
	"unity.service/pkg/org.unity.sdk/logx"
	"unity.service/pkg/org.unity.sdk/setting"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func init() {
	setting.Setup()
	logx.Setup()
	runtime.GOMAXPROCS(-1)
}



// @title ISS
// @version 2.0
// @description ISS
// @termsOfService  https://paymedia.yuque.com/os/hb9uhc/vgkey3#k5NZw
// @contact.name org.pm.iss.service
// @contact.url

// @license.name GIN.swagger
// @license.url  https://github.com/swaggo/swag/blob/master/README_zh-CN.md#%E5%A6%82%E4%BD%95%E4%B8%8Egin%E9%9B%86%E6%88%90
// @host https://3platz.t.dev.pay.fun/
// @BasePath /

func main() {
	//api.Es = eventsource.New(nil, nil)
	//defer api.Es.Close()

	//go run ariga.io/entimport/cmd/entimport -dsn "mysql://root:root@tcp(127.0.0.1:3306)/unity" -tables "users"
	cli.Execute()
}
