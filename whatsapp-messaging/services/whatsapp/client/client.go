package client

import (
	"go.mau.fi/whatsmeow"
)

var WhatsappClients map[string]*whatsmeow.Client


func init(){
	WhatsappClients = make(map[string]*whatsmeow.Client, 0)
}