package main

import (
	"log"
	"whatsapp-messaging/internal/boot"
)

func main() {
	if err := boot.Boot(); err != nil {
		log.Panic(err)
	}
}
