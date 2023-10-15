package services

import (
	"context"
	"whatsapp-messaging/internal/logger"
)

type ServiceIniterFunc func() error

var ServiceIniters []ServiceIniterFunc

func InitServices() error {
	logger.Info(context.Background(), "Initiating Services...")
	for _, service := range ServiceIniters {
		if err := service(); err != nil {
			return err
		}
	}

	return nil
}

func RegisterService(service ServiceIniterFunc) {
	ServiceIniters = append(ServiceIniters, service)
}
