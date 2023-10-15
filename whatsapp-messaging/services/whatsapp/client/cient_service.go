package client

import (
	"context"
)

func Login(ctx context.Context) (msg, qr string, timeout int, err error) {
	InitClientWithNewDevice(ctx)

	qrResponse, qrTimeout, err := LoginClient(ctx)
	if err != nil {
		return
	}

	if qrResponse == "" {
		return "reconnected successfully", "", 0, nil
	}

	return "QR Code genereated successfully", qrResponse, qrTimeout, nil
}
