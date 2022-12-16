package cli

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unity.service/pkg/org.unity.sdk/logx"
	"unity.service/pkg/org.unity.sdk/setting"
	"unity.service/routers"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "cli tool",
	Long:  `cli tool`,
	Run: func(cmd *cobra.Command, args []string) {
		runCron()
		runHttp()
	},
}



/**
 * runHttp
 * @Description:
 */
func runHttp() {
	//fixme eros
	//go example.Work()

	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	go func() {
		_ = server.ListenAndServe()
	}()
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  	<-quit
	log.Println(time.Now().Format("2006-01-02 15:04:05")+"Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 3 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}

}


// 定时任务
func runCron() {

	logx.G(nil).Info(fmt.Sprintf(`---cron: Sale:%s , Merchant:%s ,  Store:%s`, setting.Cron.TaskSalesDayCron, setting.Cron.TaskMerchantCron, setting.Cron.TaskStoreCron), func() string {return `iss.cli.cron`})
}
