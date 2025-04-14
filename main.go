package main

import (
	"kvstore/channels"
	"kvstore/db"
	"kvstore/http"
	_ "net/http/pprof" // Import pprof for profiling
)

func main() {

	serverStarted := make(chan struct{}) // Channel to signal when the server is ready
	done := make(chan bool, 1)

	go channels.Requests()
	go http.StartServer(serverStarted, done)

	<-serverStarted
	go db.Seed(100)

	<-done
}
