package middleware

import (
	"fmt"
	"runtime"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/server/web/context"
)
func RecoverPanic(ctx *context.Context, config *web.Config) {
    if err := recover(); err != nil {
        ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
        var stack []string
        for i := 1; ; i++ {
            _, file, line, ok := runtime.Caller(i)
            if !ok {
                break
            }
            logs.Critical(fmt.Sprintf("%s:%d", file, line))
            stack = append(stack, fmt.Sprintf("%s:%d \n", file, line))
        }
        //Display error
        data := map[string]interface{}{
            "ret":           4000,
            "AppError":      fmt.Sprintf("%v", err),
            "RequestMethod": ctx.Input.Method(),
            "RequestURL":    ctx.Input.URI(),
            "RemoteAddr":    ctx.Input.IP(),
            "Stack":         stack,
            //"BeegoVersion":  beego.VERSION,
            "GoVersion": runtime.Version(),
        }
        _ = ctx.Output.JSON(data, true, true)
        if ctx.Output.Status != 0 {
            ctx.ResponseWriter.WriteHeader(ctx.Output.Status)
        } else {
            ctx.ResponseWriter.WriteHeader(500)
        }
    }
}