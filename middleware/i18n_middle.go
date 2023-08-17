package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"
)

// I18nMiddleware is the middleware layer to translate all the HTTP requests - NOT USED
func I18nMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := httptest.NewRecorder()

		next.ServeHTTP(rec, r)

		fmt.Println(">>>>>>>>>>>>>>>>> I18n middleware <<<<<<<<<<<<<<<<<")
		ck := r.Cookies()
		fmt.Println("Cookies : ")
		for _, v := range ck {
			fmt.Printf("%s: %v \n", v.Name, v.Value)
			//ctx.ResponseWriter.Header()[k] = v
		}
		fmt.Println("Headers : ")
		for k, v := range w.Header() {
			fmt.Printf("%s: %v \n", k, v)
			//ctx.ResponseWriter.Header()[k] = v
		}
		if strings.Contains(rec.Header().Get("Content-Type"), "text/html") {
			start := time.Now()

			// we copy the original headers first - headers write out after w.WriteHeader(200) bellow
			for k, v := range rec.Header() {
				fmt.Printf("%s: %v \n", k, v)
				w.Header()[k] = v
			}
			w.Header().Set("X-Localised", "yes") //just for fun

			//original body length
			clen := len(rec.Body.Bytes())

			// The body hasn't been written (to the real RW) yet,
			// so we can prepend some data.
			//data := []byte("Middleware says hello again. ")
			data := "Middleware says hello again. "
			clen += len(data)
			w.Header().Set("Content-Length", strconv.Itoa(clen))

			// only then the status code, as this call writes out the headers
			w.WriteHeader(200)
			// finally, write out our data

			s := rec.Body.String()
			s = s + data
			//			fmt.Printf("Body length : %d \n", len(s))
			//_, err := w.Write(rec.Body.Bytes())
			_, err := w.Write([]byte(s))
			if err != nil {
				fmt.Printf(" Error %v", err.Error())
			}
			fmt.Printf("[Execution time: %v] \n", time.Since(start))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// ResponseWriterWrapper struct is used to log the response
type ResponseWriterWrapper struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rww *ResponseWriterWrapper) Write(p []byte) (int, error) {
	rww.body.Write(p)
	return rww.ResponseWriter.Write(p)
}

// func (rww ResponseWriterWrapper) Write(buf []byte) (int, error) {
// 	rww.body.Write(buf)
// 	return (*rww.w).Write(buf)
// }

// Header function overwrites the http.ResponseWriter Header() function
func (rww *ResponseWriterWrapper) Header() http.Header {
	return rww.ResponseWriter.Header()

}

// WriteHeader function overwrites the http.ResponseWriter WriteHeader() function
func (rww *ResponseWriterWrapper) WriteHeader(status int) {
	rww.statusCode = status
	rww.ResponseWriter.WriteHeader(status)
}

// func (rww ResponseWriterWrapper) String() string {
// 	var buf bytes.Buffer

// 	buf.WriteString("Response:")

// 	buf.WriteString("Headers:")
// 	for k, v := range (*rww.w).Header() {
// 		buf.WriteString(fmt.Sprintf("%s: %v", k, v))
// 	}

// 	buf.WriteString(fmt.Sprintf(" Status Code: %d", *(rww.statusCode)))

// 	buf.WriteString("Body")
// 	buf.WriteString(rww.body.String())
// 	return buf.String()
// }
