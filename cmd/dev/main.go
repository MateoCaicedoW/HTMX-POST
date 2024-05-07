package main

import (
	"blogpost/internal"
	"fmt"

	"github.com/leapkit/core/gloves"
)

func main() {
	fmt.Println("Starting the application in development mode")
	err := gloves.Start(
		"cmd/app/main.go",

		internal.GlovesOptions...,
	)

	if err != nil {
		fmt.Println(err)
	}
}
