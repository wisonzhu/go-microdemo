package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"gomicrodemo/helloworld/handler"
	pb "gomicrodemo/helloworld/proto"
)

const (
	ServerName = "go.micro.srv.HelloWorld" // server name
)

func main() {
	// Create service
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Version("latest"),
	)

	// Register handler
	if err := pb.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
