package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func (s *Server) Address() string {
	return s.Addr
}

func (s *Server) IsAlive() bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(s.Addr)
	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	s.Proxy.ServeHTTP(w, r)
}

type Server struct {
	Addr  string
	Proxy httputil.ReverseProxy
}

func NewServer(addr string) Server {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatal("failed to parse error")
	}
	return Server{
		Addr:  addr,
		Proxy: *httputil.NewSingleHostReverseProxy(u),
	}
}
