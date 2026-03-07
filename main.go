package main

import (
	"fmt"
	"log"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.SetUser("Gallus"); err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)

}
