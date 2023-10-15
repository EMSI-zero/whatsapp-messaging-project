package auth

import (
	"context"
	"net/http"
	"strconv"
	"whatsapp-messaging/internal/contextmanager"
	"whatsapp-messaging/services/whatsapp/user"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) (context.Context, error) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	jid, err := user.GetJID(c.Request.Context(), int64(userID))
	if err != nil {
		return nil, err
	}

	ctx, err := contextmanager.NewUserContext(c.Request.Context(), int64(userID), jid)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}
