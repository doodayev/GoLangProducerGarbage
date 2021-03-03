module go-producer-consumer-v1-main

go 1.15

replace github.com/my/repo => ../

require (
	github.com/go-redis/redis/v8 v8.5.0
	github.com/gorilla/mux v1.8.0
)
