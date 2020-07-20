package main

import (
	"log"
	"os"
)

type conf struct {
	Domain      string
	Local       bool
	VaultPrefix string
}

func getConfig() conf {
	c := conf{}
	c.Domain = os.Getenv("DOMAIN")
	c.Local = false
	c.VaultPrefix = os.Getenv("VAULT_PREFIX")
	if c.Domain == "" {
		c.Domain = "localhost"
	}
	if c.Domain == "localhost" {
		c.Local = true
	}

	log.Println("[INFO] using domain:", c.Domain, ", vault prefix:", c.VaultPrefix)

	return c
}
