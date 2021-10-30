package main

import (
	"fmt"
	. "github.com/torchcc/crank4go/connector"
	"net/http"
	"net/url"
)
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Client-VM1")
}

func main() {
	targetURI, _ := url.Parse("http://localhost:7000")
	routerURI, _ := url.Parse("wss://localhost:1323") // should be the port which your Router Registration server listens on

	connectorConfig := NewConnectorConfig2(targetURI, "my-service", []*url.URL{routerURI}, "my-service-component-name", nil).
		SetSlidingWindowSize(2)
	_ = CreateAndStartConnector(connectorConfig)

	http.HandleFunc("/my-service", HelloHandler)
	http.ListenAndServe(":7000", nil)
	// and then you can query your api gateway to access your service. 
	// e.g. if your router listens on https://localhost:1323, then you can access  https://localhost:1323/my-service
}
