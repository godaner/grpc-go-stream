package main

import (
	"grpc-go-stream/client"
	"time"
	"grpc-go-stream/server"
)

func main(){
	go func() {
		server.Run()
	}()
	time.Sleep(3*time.Millisecond)
	client.Run()
}
