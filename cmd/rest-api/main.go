package main

import (
	"fmt"
	"os"

	rest_api "github.com/brunoeduardodev/trellogo/internal/infra/rest-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	fmt.Println("Hello world")

	rest_api.Start(os.Getenv("PORT"))
}
