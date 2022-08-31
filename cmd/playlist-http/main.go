package main

import (
	"log"

	"github.com/nabiel-syarif/playlist-api/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to init conig, err : %v", err)
	}

	err = startApp(cfg)

	if err != nil {
		log.Fatalf("something went wrong when running app. err : %v", err)
	}
}
