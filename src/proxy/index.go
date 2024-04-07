package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy() {

	remote, err := url.Parse("http://google.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("remote: ", remote)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	// use http.Handle instead of http.HandleFunc when your struct implements http.Handler interface
	http.Handle("/", &Handler{proxy, remote})
	fmt.Println("Listening on port 3000")
	err = http.ListenAndServe("0.0.0.0:3000", nil)
	if err != nil {
		panic(err)
	}
}

type Handler struct {
	proxy          *httputil.ReverseProxy
	destinationUrl *url.URL
}

func (ph *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Host = ph.destinationUrl.Host
	r.URL.Host = ph.destinationUrl.Host
	r.URL.Scheme = ph.destinationUrl.Scheme
	r.RequestURI = ""
	ph.proxy.ServeHTTP(w, r)
}
