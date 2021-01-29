package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/keuller/simple-api/internal/application/web"
)

func main() {
	_ = godotenv.Load()

	http_port := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	server := web.Server()

	log.Println(fmt.Sprintf("Service up and running on http://localhost%s", http_port))
	server.Run(http_port)
}
