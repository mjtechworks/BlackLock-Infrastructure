package main

import (
	"net/http"

	"github.com/koding/tunnel"
)

func main() {
	cfg := &tunnel.ServerConfig{}
	server, _ := tunnel.NewServer(cfg)
	server.AddHost("https://dlm-investments0.webnode.com", "1234")
	http.ListenAndServe(":80", server)
}
