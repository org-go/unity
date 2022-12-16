package logx

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gookit/color"
    "strings"
    "time"
    "unity.service/pkg/org.unity.sdk/setting"
)


var C = getLoggerWitchGinContext
var G = getLoggerWithContext



type (
    loggers struct{
        c *gin.Context
        ctx *context.Context
        level int
        color *color.Theme
    }
    CallError func() string
)

func getLoggerWithContext(ctx *context.Context) *loggers {
    return &loggers{ctx: ctx}
}

func getLoggerWitchGinContext(ctx *gin.Context) *loggers {
   // pkg.Trace(ctx).AddEvent(ctx.FullPath(), nil)
    return &loggers{c: ctx}
}

/**
 * $(V)
 * @Description:
 * @receiver l
 * @param level
 * @return *logger
 */
func (l *loggers) V (level int) *loggers {
    l.level = level
    return l
}



/**
 * $(Error)
 * @Description:
 * @receiver l
 * @param err
 * @param callErrors
 */
func (l loggers) Error (err error, callErrors ...CallError)  {
    l.color = color.Error
    str := fmt.Sprintf(`{LEVEL}:{%d} --- [ERROR] %s ### %s`, l.level, callError(callErrors...), err.Error())
    fmt.Println(str)
    write(str)
}


/**
 * $(Info)
 * @Description:
 * @receiver logger
 * @param str
 * @param callErrors
 */
func (l loggers) Info (str string, callErrors ...CallError)  {

    l.color = color.Info
    str = fmt.Sprintf("{LEVEL}:{%d} --- [LOG] %v ### %s \n", l.level, callError(callErrors...), str)
    fmt.Println(str)
    if setting.ServerSetting.RunMode == `debug` {
        write(str)
    }

}

/**
 * $(Debug)
 * @Description:
 * @receiver logger
 * @param any
 * @param callErrors
 */
func (l loggers) Debug (any interface{}, callErrors ...CallError)  {

    l.color = color.Debug
    a, _ := json.Marshal(any)
    str := fmt.Sprintf("{LEVEL}:{%d} --- [DEBUG] %s ### %v \n", l.level, callError(callErrors...), string(a))
    fmt.Println(str)
    if setting.ServerSetting.RunMode == `debug` {
        write(str)
    }
}

/**
 * $(callError)
 * @Description:
 * @param callErrors
 * @return string
 */
func callError(callErrors ...CallError) string {

    var str []string
    for _, callError := range callErrors {
        str = append(str, callError())
    }
    return strings.Join(str, "\n")

}

func write(str string)  {
    _, _ = F.WriteString(fmt.Sprintf(`%s = %s`,time.Now().Format(`2006-01-02 15:04:05`), str)+"\n")
}
