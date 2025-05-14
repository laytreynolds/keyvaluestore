package main

import (
	"kvstore/channels"
	"kvstore/http"
	"kvstore/store"
	_ "net/http/pprof" // Import pprof for profiling
)

func main() {

	serverStarted := make(chan struct{}) // Channel to signal when the server is ready
	done := make(chan bool, 1)           // Block the main goroutine until the server has shutdown

	go channels.Requests()
	go http.StartServer(serverStarted, done)

	<-serverStarted

	store.Store.InitData()

	<-done
}
