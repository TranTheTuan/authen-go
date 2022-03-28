package pubsub

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func EventDispatcherMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqBodyBytes []byte
		if r.Body != nil {
			reqBodyBytes, _ = ioutil.ReadAll(r.Body)
			// Restore the io.ReadCloser to its original state
			r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBodyBytes))
		}
		next.ServeHTTP(w, r)
		if r.Method != "GET" && !strings.Contains(r.RequestURI, "login") {
			entry := log.WithFields(log.Fields{
				"request_body": string(reqBodyBytes),
			})
			brokerW.DispatchEvent(entry)
		}
	})
}
