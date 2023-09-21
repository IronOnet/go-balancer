package main 

import (
	"fmt" 
	"log" 
	"net/http" 
	"net/http/httputil" 
	"net/url"
) 

var servers = []string{
	"http://localhost:8080", 
	"http://localhost:8081", 
	"http://locahost:8082", 
} 

func main(){
	balancer := NewLoadBalancer(servers) 
	log.Fatal(http.ListenAndServe(":80", balancer)) 
} 

type LoadBalancer struct{
	servers []*url.URL 
	index int 
} 

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// Round-robin selection of the next server 
	serverURL := lb.servers[lb.index] 
	lb.index = (lb.index + 1) % len(lb.servers) 

	proxy := httputil.NewSingleHostReverseProxy(serverURL) 
	proxy.ServeHTTP(w, r) 
}
