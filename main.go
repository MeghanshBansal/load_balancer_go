package main

import (
	"log"
	"net/http"
	"github.com/MeghanshBansal/loadBalancerGo/loadBalancer"
	"github.com/MeghanshBansal/loadBalancerGo/server"
)


func main() {
	servers := []server.Server{
		server.NewServer("https://www.facebook.com"),
		server.NewServer("https://www.duckduckgo.com"),
		server.NewServer("https://www.google.com"),
	}

	lb := loadbalancer.NewLoadBalancer(":8000", servers)
	handlerRedirct := func(w http.ResponseWriter, r *http.Request) {
		lb.ServeProxy(w, r)
	}

	http.HandleFunc("/", handlerRedirct)
	http.ListenAndServe(lb.Port, nil)
	log.Println("Load Balancing on localhost", lb.Port)
}
