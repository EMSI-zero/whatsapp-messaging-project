package user

import (
	"context"
	"fmt"
	"whatsapp-messaging/internal/logger"
	"whatsapp-messaging/internal/repository/dbrepo"
)

func SetJID(ctx context.Context, jid string) error {
	db, err := dbrepo.GetGormConn(ctx)
	if err != nil {
		return err
	}

	userId, _, err := ReadUserContext(ctx)
	if err != nil {
		return err
	}

	if err := db.Create(&User{UserID: userId, Jid: jid}).Error; err != nil {
		return err
	}

	return nil
}

type UserContextKey struct{}
type JIDContextKey struct {}

func NewUserContext(ctx context.Context, userID int64, jid string) (context.Context, error) {
	if userID == 0 {
		err := fmt.Errorf("no user id found")
		logger.Error(ctx, err)
		return nil, err
	}

	ctx = context.WithValue(ctx, &UserContextKey{}, userID)
	if jid != "" {
		ctx = context.WithValue(ctx, &JIDContextKey{}, jid)
	}

	return ctx, nil
}

func ReadUserContext(ctx context.Context) (userId int64, jid string, err error) {
	userIdValue := ctx.Value("user_id")
	if userIdValue == nil {
		return 0, "", fmt.Errorf("no user id found")
	}
	jidValue := ctx.Value("jid")

	return userIdValue.(int64), jidValue.(string), nil
}
