package setting

import (
    "fmt"
    "github.com/go-ini/ini"
    "log"
    "os"
    "strings"
    "time"
    "unity.service/apps/v1/entity/platform"
    "unity.service/apps/v1/entity/platform/coreSystem"
    "unity.service/apps/v1/repository/repo"
)

var (

    // app
    SdkSetting = &sdk{}

    // iss
    issSetting    = &iss{}
    IssFtpSetting = &issFtp{}
    // core
    coreSetting    = &core{}
    CoreFtpSetting = &coreFtp{}

    // cps
    CpsSetting    = &cps{}
    CpsFtpSetting = &cpsFtp{}

    // env
    cfg           *ini.File
    mysqlDB       = &repo.DataBase{}
    RedisConf     = &redisConf{}
    AppSetting    = &app{}
    ServerSetting = &server{}

    //cron
    Cron    = &cron{}
)

type (
    redisConf struct {
        Host     string
        Password string
        Db       int
    }

    app struct {
        AppName            string
        RuntimeRootPath    string
        LogSavePath        string
        LogSaveName        string
        LogFileExt         string
        TimeFormat         string
        ImSalesSavePath    string
        ImMerchantSavePath string
        ImStoreSavePath    string
        ImTmpSavePath      string
        Env                string
    }

    server struct {
        RunMode      string
        HttpPort     int
        ReadTimeout  time.Duration
        WriteTimeout time.Duration
    }

    core struct {
        Host string
        Url  string
    }

    sdk struct {
        Sign              string
        CodeTimeoutSecond int
        CodeValidDay      int
        GpsSafeDistance   int
    }

    cps struct {
        Sign string
    }

    iss struct {

    }

    issFtp struct {
        Host     string
        Port     int
        User     string
        Password string
        Path     string
    }

    cpsFtp struct {
        Host     string
        Port     int
        User     string
        Password string
        Path     string
    }
    coreFtp struct {
        Host     string
        Port     int
        User     string
        Password string
        Path     string
    }

    cron struct {
        TaskSalesDayCron string
        TaskMerchantCron string
        TaskStoreCron string
    }
)

// Setup initialize the configuration instance
func Setup() {
    var err error
    env := os.Getenv("GOENV")
    switch env {
    case `local`, `dev`:
        AppSetting.Env = `pm_` + env
    case `iss`, `uat`, `pro`:
        AppSetting.Env = env
    default:
        AppSetting.Env = `dev`
    }
    cfg, err = ini.Load("conf/" + AppSetting.Env + ".ini")
    if err != nil {
        log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
    } else {
        config()
    }
    repo.Registry(*mysqlDB)
    platform.Core = coreSystem.RegistryCorePlatform(&coreSystem.CoreConfig{
        Appid:     "",
        Appsecret: "",
        Site:      coreSetting.Host,
        Url:       coreSetting.Url,
    })
    echo()

}

// mapTo map section
func mapTo(section string, v interface{}) {
    err := cfg.Section(section).MapTo(v)
    if err != nil {
        log.Fatalf("Cfg.MapTo %s err: %v", section, err)
    }
}

func config() {
    mapTo("app", AppSetting)
    mapTo("server", ServerSetting)
    mapTo("sdk", SdkSetting)
    mapTo("core", coreSetting)
    mapTo("core.ftp", CoreFtpSetting)
    mapTo("cps", CpsSetting)
    mapTo("cps.ftp", CpsFtpSetting)
    mapTo("db", mysqlDB)
    mapTo("redis", RedisConf)
    mapTo("cron", Cron)

    ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
    ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

}

func echo() {
    fmt.Println(strings.Repeat(`#`, 100))
    fmt.Println(strings.Repeat(`#`, 44) + ` {ISS STATUS} ` + strings.Repeat(`#`, 42))
    fmt.Println(strings.Repeat(`#`, 100))
    fmt.Println(`[project]  #`, AppSetting.AppName)
    fmt.Println(`[port]     #`, ServerSetting.HttpPort)
    fmt.Println(`[env]      #`, AppSetting.Env)
    fmt.Println(`[mysql]    #`, mysqlDB)
    fmt.Println(`[redis]    #`, RedisConf)
    fmt.Println(`[sdk]      #`, SdkSetting)
    fmt.Println(`[cps]      #`, CpsSetting)
    fmt.Println(`[cps.ftp]  #`, CpsFtpSetting)
    fmt.Println(`[core]     #`, coreSetting)
    fmt.Println(`[core.ftp] #`, CoreFtpSetting)
    fmt.Println(`[cron]     #`, Cron)

    fmt.Println(strings.Repeat(`#`, 100))
    fmt.Printf("\n\n")
}
