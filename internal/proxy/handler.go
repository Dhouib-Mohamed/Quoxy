package proxy

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/util/log"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Handler struct {
	proxy          *httputil.ReverseProxy
	destinationUrl *url.URL
}

func (ph *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debug(fmt.Sprintf("Proxying request from %s to %s", r.RemoteAddr, r.URL.String()))
	token := r.Header.Get("Proxy-Authorization")
	if token == "" {
		http.Error(w, "Token is not provided", http.StatusUnauthorized)
		log.Error(fmt.Errorf("token in request %s is not provided , please provide a correct one in the Proxy-Authorization Header", r.URL.String()))
		return
	}
	t := database.Token{}
	err := t.Use(token)
	if err != nil {
		code, msg := err.GetError()
		log.Error(fmt.Errorf("error using token: %s", msg))
		http.Error(w, msg, code)
		return
	}
	r.Host = ph.destinationUrl.Host
	r.URL.Host = ph.destinationUrl.Host
	r.URL.Scheme = ph.destinationUrl.Scheme
	r.RequestURI = ""
	log.Debug(fmt.Sprintf("Proxying request from %s to %s succeded", r.RemoteAddr, r.URL.String()))
	ph.proxy.ServeHTTP(w, r)
}
