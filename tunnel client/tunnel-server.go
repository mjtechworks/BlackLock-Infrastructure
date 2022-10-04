package main

import (
	"net/http"

	"github.com/koding/tunnel"
)

func main() {
	cfg := &tunnel.ServerConfig{}
	server, _ := tunnel.NewServer(cfg)
	server.AddHost("http://dlm-investments.bitrix24.site", "1234")
	http.ListenAndServe(":80", server)
}
