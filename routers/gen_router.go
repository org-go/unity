package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1671084367)
	ginrpc.AddGenOne("WelfareApi.List", "/welfare/list", []string{"POST"})
	ginrpc.AddGenOne("WelfareClassifyApi.Option", "/welfare_classify/option", []string{"GET"})
	ginrpc.AddGenOne("RankApi.Option", "/rank/option", []string{"GET"})
	ginrpc.AddGenOne("RightsClassifyApi.Option", "/rights_classify/option", []string{"GET"})
	ginrpc.AddGenOne("PlatformApi.Option", "/platform/option", []string{"GET"})
	ginrpc.AddGenOne("RankApi.Option", "/rank/option", []string{"GET"})
}
