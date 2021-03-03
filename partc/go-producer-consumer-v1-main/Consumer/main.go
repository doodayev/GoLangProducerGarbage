package main

import (
	svc "go-producer-consumer-v1-main/Consumer/service"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
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
	port := ":8080"
	producer := "http://localhost:9090"
	consumer := svc.ConsumerService{
		StopChan: make(chan bool),
		Producer: producer,
		Redis: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
	}
	consumer.Initialize()
	HttpKeepAlive(port)
}
