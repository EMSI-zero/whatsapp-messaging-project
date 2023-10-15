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
	db := dbrepo.GetDBConnPool()
	logger.Info(context.Background(), db)
	dataStoreContainer = sqlstore.NewWithDB(db, "postgres", nil)
	
	logger.Info(context.Background(), "upgrading database migrations...")
	logger.Info(context.Background(), dataStoreContainer)
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
