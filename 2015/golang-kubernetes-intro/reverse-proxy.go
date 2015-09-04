package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("http://www.bite-code.com/")
	if err != nil {
		log.Fatal("Error parsing the proxy url")
	}
	p := httputil.NewSingleHostReverseProxy(target)
	http.Handle("/", p)
	http.ListenAndServe(":8080", nil)
}
