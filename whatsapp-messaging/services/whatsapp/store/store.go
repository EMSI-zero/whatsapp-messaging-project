package store

import (
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
	err := dataStoreContainer.Upgrade()
	if err != nil {
		return err
	}
	return nil
}

func GetDataStore() *sqlstore.Container {
	return dataStoreContainer
}
