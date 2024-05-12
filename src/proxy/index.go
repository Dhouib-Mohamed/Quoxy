package proxy

import (
	"api-authenticator-proxy/src/utils/log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy() {

	remote, err := url.Parse("http://google.com")
	log.Fatal(err)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.Handle("/", &Handler{proxy, remote})
	log.Info("Listening on port 3000")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}
