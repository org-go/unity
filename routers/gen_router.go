package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1672040692)
	ginrpc.AddGenOne("FeedbackApi.Create", "/feedback/create", []string{"POST"})
	ginrpc.AddGenOne("FeedbackApi.Page", "/feedback/page", []string{"POST"})
	ginrpc.AddGenOne("PlatformApi.Option", "/platform/option", []string{"GET"})
	ginrpc.AddGenOne("RankApi.Option", "/rank/option", []string{"GET"})
	ginrpc.AddGenOne("RightsClassifyApi.Option", "/rights_classify/option", []string{"GET"})
	ginrpc.AddGenOne("RightsApi.Page", "/rights/page", []string{"POST"})
	ginrpc.AddGenOne("WelfareClassifyApi.Option", "/welfare_classify/option", []string{"GET"})
	ginrpc.AddGenOne("WelfareApi.Page", "/welfare/page", []string{"POST"})
	ginrpc.AddGenOne("TaskClassifyApi.Option", "/task_classify/option", []string{"GET"})
	ginrpc.AddGenOne("TaskApi.Create", "/task/create", []string{"POST"})
	ginrpc.AddGenOne("TaskApi.Page", "/task/page", []string{"POST"})
	ginrpc.AddGenOne("TaskApi.Update", "/task/update", []string{"POST"})
	ginrpc.AddGenOne("TaskOrderApi.Receive", "/task_order/receive", []string{"GET"})
}
