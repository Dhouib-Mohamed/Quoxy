package proxy

import (
	"api-authenticator-proxy/util/config"
	"api-authenticator-proxy/util/log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy() {
	if !config.GetIsProxyEnabled() {
		return
	}
	port := config.GetProxyPort()
	target := config.GetProxyTarget()
	remote, err := url.Parse(target)
	log.Fatal(err)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.Handle("/", &Handler{proxy, remote})
	log.Info("Proxy running on port:", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
