package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/server/web"

	context2 "github.com/beego/beego/v2/server/web/context"
)

// type MyResponseWriter struct {
// 	http.ResponseWriter
// 	buf *bytes.Buffer
// }

// func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
// 	return mrw.buf.Write(p)
// }

func TestMiddle() {

	var filter = func(ctx *context2.Context) {
		fmt.Println(">>>>>>>>>>>>>>>>> Test middleware <<<<<<<<<<<<<<<<<")
		for k, v := range ctx.ResponseWriter.Header() {
			fmt.Printf("%s: %v \n", k, v)
			//ctx.ResponseWriter.Header()[k] = v
		}
		if strings.Contains(ctx.ResponseWriter.Header().Get("Content-Type"), "text/html") {

			url := strings.TrimLeft(ctx.Input.URL(), "/")
			fmt.Printf(">>>>>>>>>>>>>>>> URL :%s \n", url)
			// mrw := &MyResponseWriter{
			// 	ResponseWriter: ctx.ResponseWriter,
			// 	buf:            &bytes.Buffer{},
			// }

			w := ctx.ResponseWriter.ResponseWriter
			w.Header().Set("X-Localised-This", "yes") //just for fun
			//w.WriteHeader(403)
			data := "403 Forbidden\n"
			w.Header().Set("Content-Length", strconv.Itoa(6577+len(data)))
			// _, err := w.Write([]byte(data))
			err := ctx.Output.Body([]byte(data))
			if err != nil {
				fmt.Printf("Error : %v", err.Error())
			}

			return
		}
	}
	web.InsertFilter("/*", web.FinishRouter, filter)
}
