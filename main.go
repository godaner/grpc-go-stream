package main

import (
	"grpc-go-stream/server"
	"grpc-go-stream/client"
)

func main(){
	server.Run()
	client.Run()
}
