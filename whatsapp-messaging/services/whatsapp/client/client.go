package client

import (
	"context"
	"encoding/base64"
	"fmt"
	"whatsapp-messaging/internal/contextmanager"
	"whatsapp-messaging/internal/logger"
	"whatsapp-messaging/services/whatsapp/store"

	qrCode "github.com/skip2/go-qrcode"

	"go.mau.fi/whatsmeow"
	wmstore "go.mau.fi/whatsmeow/store"
)

var WhatsappClients map[string]*whatsmeow.Client

func init() {
	WhatsappClients = make(map[string]*whatsmeow.Client, 0)
}

func InitClientWithNewDevice(ctx context.Context) error {
	device := store.GetDataStore().NewDevice()
	return InitClient(ctx, device)
}

func InitClient(ctx context.Context, device *wmstore.Device) error {
	jid := ctx.Value(contextmanager.JIDContextKey{}).(string)
	if jid == "" {
		err := fmt.Errorf("no jid found")
		logger.Error(ctx, err)
		return err
	}

	if WhatsappClients[jid] != nil {
		return nil
	}

	WhatsappClients[jid] = whatsmeow.NewClient(device, nil)
	WhatsappClients[jid].EnableAutoReconnect = true
	WhatsappClients[jid].AutoTrustIdentity = true

	return nil
}

func WhatsAppGenerateQR(qrChan <-chan whatsmeow.QRChannelItem) (string, int) {
	qrChanCode := make(chan string)
	qrChanTimeout := make(chan int)

	go func() {
		for evt := range qrChan {
			if evt.Event == "code" {
				qrChanCode <- evt.Code
				qrChanTimeout <- int(evt.Timeout.Seconds())
			}
		}
	}()

	qrTemp := <-qrChanCode
	qrPNG, _ := qrCode.Encode(qrTemp, qrCode.Medium, 256)

	return base64.StdEncoding.EncodeToString(qrPNG), <-qrChanTimeout
}
