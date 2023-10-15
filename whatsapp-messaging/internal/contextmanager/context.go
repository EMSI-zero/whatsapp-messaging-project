package contextmanager

import (
	"context"
	"fmt"
)

type UserContextKey struct{}
type JIDContextKey struct{}

func NewUserContext(ctx context.Context, userID int64, jid string) (context.Context, error) {
	if userID == 0 {
		err := fmt.Errorf("no user id found")
		return nil, err
	}

	ctx = context.WithValue(ctx, UserContextKey{}, userID)
	if jid != "" {
		ctx = context.WithValue(ctx, JIDContextKey{}, jid)
	}

	return ctx, nil
}

func ReadUserContext(ctx context.Context) (userId int64, jid string, err error) {
	userIdValue := ctx.Value(UserContextKey{})
	if userIdValue == nil {
		return 0, "", fmt.Errorf("no user id found")
	}
	jidValue := ctx.Value(JIDContextKey{})

	return userIdValue.(int64), jidValue.(string), nil
}
