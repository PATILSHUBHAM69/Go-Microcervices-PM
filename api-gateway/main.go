package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var servicePorts = map[string]string{
	"service1": "8080",
	"service2": "8081",
	"service3": "8082",
	"service4": "8083",
}

func reverseProxy(service string) http.Handler {
	targetURL, err := url.Parse("http://localhost:" + servicePorts[service])
	if err != nil {
		log.Fatal("Error parsing target URL:", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	return proxy
}

func staticHandler() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func main() {
	http.Handle("/service1/", reverseProxy("service1"))
	http.Handle("/service2/", reverseProxy("service2"))
	http.Handle("/service3/", reverseProxy("service3"))
	http.Handle("/service4/", reverseProxy("service4"))

	http.Handle("/static/", http.StripPrefix("/static/", staticHandler()))

	fmt.Println("API Gateway is running on port 8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
