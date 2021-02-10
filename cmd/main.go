package main

import (
	"fmt"
	"log"

	"github.com/keuller/simple-api/internal/application/web"
	"github.com/keuller/simple-api/internal/config"
)

func main() {
	config := config.Get()

	httpPort := fmt.Sprintf(":%s", config.GetString("HTTP_PORT"))

	server := web.Server()

	log.Println(fmt.Sprintf("Service up and running on http://localhost%s", httpPort))
	server.Run(httpPort)
}
