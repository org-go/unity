package routers

import (
	//"github.com/arl/statsviz"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	//ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xxjwxc/ginrpc"
	"net/http"
	"unity.service/apps/v1/api"
	_ "unity.service/docs"
	"unity.service/middleware"
	//"github.com/gin-contrib/pprof"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.Cors())
	router.Use(gin.ErrorLogger())
	router.Use(ginprom.PromMiddleware(nil))
	router.Use(gin.Recovery())
	auc := router.Group(``)
	auc.StaticFS(`/logs`, http.Dir(`./logs`))
	//auc.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, ``))
	auc.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	v := new(version).option(auc)
	v.http()

	return router
}

type version struct {
	v   string
	auc *gin.RouterGroup
}

func (self version) option(auc *gin.RouterGroup) *version {
	self.v = `v1`
	self.auc = auc
	return &self
}

func (self version) http() {

	base := ginrpc.New()
	base.Register(self.auc, new(api.WelfareApi))
	base.Register(self.auc, new(api.WelfareClassifyApi))
	base.Register(self.auc, new(api.RankApi))
	base.Register(self.auc, new(api.RightsClassifyApi))
	base.Register(self.auc, new(api.PlatformApi))
	base.Register(self.auc, new(api.RankApi))

}
