package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/studyzy/simpleJsonrpc/ethapi"
)

const (
	DefaultHTTPHost = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort = 8545        // Default TCP port for the HTTP RPC server
)

func main() {
	b := &MockBackend{}
	rpcAPIs := ethapi.GetAPIs(b)
	glogger := log.NewLogger(log.NewTerminalHandler(os.Stdout, false))
	config := &node.Config{
		HTTPHost: DefaultHTTPHost,
		HTTPPort: DefaultHTTPPort,
		Logger:   glogger,
	}

	api, err := node.New(config)
	if err != nil {
		panic(err.Error())
	}
	api.RegisterAPIs(rpcAPIs)
	err = api.Start()
	if err != nil {
		panic(err.Error())

	}
	fmt.Println("API server started on", api.HTTPEndpoint())
	// Wait for interrupt signal to gracefully shutdown the server.
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
}
