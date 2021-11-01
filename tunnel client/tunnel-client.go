package main

import "github.com/koding/tunnel"

func main() {
	cfg := &tunnel.ClientConfig{
		Identifier: "1234",
		ServerAddr: "192.168.1.158:80",
	}

	client, err := tunnel.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	client.Start()
}
