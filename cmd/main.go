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

	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	server := web.NewServer(httpPort)

	log.Println(fmt.Sprintf("Service up and running on http://localhost%s", httpPort))
	server.Start()
}
