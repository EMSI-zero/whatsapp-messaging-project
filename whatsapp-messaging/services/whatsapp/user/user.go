package user

import (
	"context"
	"time"
	"whatsapp-messaging/internal/contextmanager"
	"whatsapp-messaging/internal/logger"
	"whatsapp-messaging/internal/repository/cacherepo"
	"whatsapp-messaging/internal/repository/dbrepo"
)

var UserJIDCache *cacherepo.DBCache[int64, string] = cacherepo.MakeCache(1*time.Minute, LoadJID)

func LoadJID(ctx context.Context, userID int64) (string, error) {
	db, err := dbrepo.GetGormConn(ctx)
	if err != nil {
		return "", err
	}

	var jid string
	if err := db.Model(&User{}).Take(&jid).Error; err != nil {
		return "", err
	}

	return jid, nil
}

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


func GetJID(ctx context.Context, userId int64) (jid string,err error){
	return UserJIDCache.Read(ctx, userId)
}