package main

import (
	"fmt"
	"log"

	"github.com/alfredoprograma/mks/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg)
}
