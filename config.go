package main

import (
	"log"
	"os"
)

type conf struct {
	Domain string
	Local  bool
}

func getConfig() conf {
	c := conf{}
	c.Domain = os.Getenv("DOMAIN")
	if c.Domain == "localhost" {
		c.Local = true
	}

	log.Println("[INFO] using domain:", c.Domain)

	return c
}
