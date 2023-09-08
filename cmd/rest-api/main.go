package main

import (
	"fmt"
	"os"

	"github.com/brunoeduardodev/trellogo/internal/infra/config"
	"github.com/brunoeduardodev/trellogo/internal/infra/database"
	"github.com/brunoeduardodev/trellogo/internal/infra/database/users"
	"github.com/brunoeduardodev/trellogo/internal/infra/rest_api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	fmt.Println("Hello world")

	config := config.NewConfig()

	fmt.Printf("PORT: %v\n", config.RestApi.Port)
	fmt.Printf("PORT: %v\n", os.Getenv("PORT"))

	db := database.NewGormDatabase()
	userRepository := users.NewUserRepository(db)

	rest_api.Start(userRepository)
}
