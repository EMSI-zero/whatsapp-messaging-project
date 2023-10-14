package user

import (
	"context"
	"whatsapp-messaging/internal/contextmanager"
	"whatsapp-messaging/internal/logger"
	"whatsapp-messaging/internal/repository/dbrepo"
)

func SetJID(ctx context.Context, jid string) error {
	db, err := dbrepo.GetGormConn(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	userId, _, err := contextmanager.ReadUserContext(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := db.Create(&User{UserID: userId, Jid: jid}).Error; err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
