package main

import (
	"log"
	"whatapp-messaging/internal/boot"
)

func main() {
	if err := boot.Boot(); err != nil {
		log.Panic(err)
	}
}