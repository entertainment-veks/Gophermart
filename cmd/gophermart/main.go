package main

import (
	"gophermart/internal/app/gophermart"
	"log"
)

func main() {
	config, err := gophermart.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := gophermart.Start(config); err != nil {
		log.Fatal(err)
	}
}
