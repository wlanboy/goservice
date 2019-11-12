package application

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

/*loggingMiddleware*/
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		} else {
			filter := strings.Replace(string(request), "\r\n", " ", -1)
			log.Println(fmt.Sprintf("%q", filter))
		}
		next.ServeHTTP(w, r)
	})
}
