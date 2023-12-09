package main

import (
	"go-web-template/server"
	"go-web-template/server/config"
	"net/http"
)

func main() {
	run()
}

func run() {
	router := http.NewServeMux()

	port := ":9999"
	db := config.CreateConnection()

	server.StartServer(router, port, db)
}
