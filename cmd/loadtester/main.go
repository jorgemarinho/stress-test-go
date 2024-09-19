package main

import (
	"github.com/jorgemarinho/stress-test-go/internal/infrastructure"
)

func main() {
	container := infrastructure.NewContainer()

	cli := container.GetCLI()

	cli.Execute()
}
