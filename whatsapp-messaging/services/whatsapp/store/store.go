package store

import (
	"context"
	"whatsapp-messaging/internal/logger"
	"whatsapp-messaging/internal/repository/dbrepo"
	"whatsapp-messaging/services"

	"go.mau.fi/whatsmeow/store/sqlstore"
)

var dataStoreContainer *sqlstore.Container

func init() {
	services.RegisterService(InitContainer)
}

func InitContainer() error {
	dataStoreContainer = sqlstore.NewWithDB(dbrepo.GetDBConnPool(), "postgres", nil)
	
	logger.Info(context.Background(), "upgrading database migrations...")
	err := dataStoreContainer.Upgrade()
	if err != nil {
		logger.Error(context.Background(), err)
		return err
	}
	logger.Info(context.Background(), "database migrations upgraded.")

	return nil
}

func GetDataStore() *sqlstore.Container {
	return dataStoreContainer
}
