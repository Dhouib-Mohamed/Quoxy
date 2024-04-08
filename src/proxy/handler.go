package proxy

import (
	"api-authenticator-proxy/src/database"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Handler struct {
	proxy          *httputil.ReverseProxy
	destinationUrl *url.URL
}

func (ph *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check if there is a provided bearer token in the request header
	token := r.Header.Get("Proxy-Authorization")
	if token == "" {
		http.Error(w, "Token is not provided", http.StatusUnauthorized)
		return
	}
	t := database.Token{}
	err := t.Use(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	r.Host = ph.destinationUrl.Host
	r.URL.Host = ph.destinationUrl.Host
	r.URL.Scheme = ph.destinationUrl.Scheme
	r.RequestURI = ""
	ph.proxy.ServeHTTP(w, r)
}
