package main

import (
	"fmt"

	rest_api "github.com/brunoeduardodev/trellogo/internal/infra/rest-api"
)

func main() {
	fmt.Println("Hello world")

	rest_api.Start()
}
