package loadbalancer

import (
	"log"
	"net/http"
	"github.com/MeghanshBansal/loadBalancerGo/server"
)

type LoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []server.Server
}

func NewLoadBalancer(port string, servers []server.Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		RoundRobinCount: 0,
		Servers:         servers,
	}
}

func (lb *LoadBalancer) getNextAvailableFunc() server.Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}

	lb.RoundRobinCount++
	return server
}

func (lb *LoadBalancer) ServeProxy(w http.ResponseWriter, r *http.Request) {
	server := lb.getNextAvailableFunc()
	log.Println("Redirecting request to server: ", server.Addr)
	server.Serve(w, r)
}