package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}
func TestMiddle2() {

	var fc = func(next web.FilterFunc) web.FilterFunc {
		return func(ctx *context.Context) {
			// do something
			fmt.Println(" >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>   FILTERCHAIN")
			for k, v := range ctx.ResponseWriter.Header() {
				fmt.Printf("%s: %v \n", k, v)
				//ctx.ResponseWriter.Header()[k] = v
			}
			url := strings.TrimLeft(ctx.Input.URL(), "/")
			fmt.Printf(">>>>>>>>>>>>>>>> URL :%s \n", url)
			mrw := &MyResponseWriter{
				ResponseWriter: ctx.ResponseWriter,
				buf:            &bytes.Buffer{},
			}

			fmt.Printf("Bytes : %v \n", len(mrw.buf.Bytes()))
			//ctx.WriteString("HAU")
			// _, err := ctx.ResponseWriter.Write([]byte("BAU"))
			// if err != nil {
			// 	fmt.Printf("Error : %v", err.Error())
			// }

			// don't forget this
			next(ctx)
			// do something
		}
	}
	web.InsertFilterChain("/*", fc)
}
