package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"gomicro/microservice"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our microservice service
	srv := microservice.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := microservice.Endpoints{
		GetEndpoint:      microservice.MakeGetEndpoint(srv),
		StatusEndpoint:   microservice.MakeStatusEndpoint(srv),
		ValidateEndpoint: microservice.MakeValidateEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("microservice is listening on port:", *httpAddr)
		handler := microservice.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}