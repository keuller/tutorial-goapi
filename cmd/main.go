package main

import (
	"fmt"
	"log"

	"github.com/keuller/simple-api/internal/application/web"
	"github.com/keuller/simple-api/internal/config"
)

func main() {
	config := config.Get()

	http_port := fmt.Sprintf(":%s", config.GetString("HTTP_PORT"))

	server := web.Server()

	log.Println(fmt.Sprintf("Service up and running on http://localhost%s", http_port))
	server.Run(http_port)
}
