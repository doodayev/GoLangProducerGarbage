package main

import (
	svc "go-producer-consumer-v1-main/Producer/service"
	"log"
	"net/http"
)

var logger *log.Logger

func HttpKeepAlive(port string) {
	errChan := make(chan error)
	go func() {
		log.Println("HTTP KeepAlive :transport", "HTTP", "started on port", port)
		errChan <- http.ListenAndServe(port, nil)
	}()
	log.Fatal("exit", <-errChan)
}

func main() {
	port := ":9090"
	consumer := "http://localhost:8080"
	producer := svc.ProducerService{
		//Remove & populate values accordingly
		StopChan: make(chan bool),
		Consumer: consumer,
	}
	producer.Initialize()
	HttpKeepAlive(port)
}
