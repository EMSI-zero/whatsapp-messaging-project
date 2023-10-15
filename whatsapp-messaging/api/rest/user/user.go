package user

import (
	"encoding/json"
	"io"
	"whatsapp-messaging/services/whatsapp/user"

	"github.com/gin-gonic/gin"
)

func AddJID(c *gin.Context) {
	ctx:= c.Request.Context()

	var req *AddJIDRequest
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		c.Error(err)
	}

	err = user.SetJID(ctx, req.JID)
	if err!= nil{
		c.Error(err)
	}
}

func AddRoutes(parent *gin.RouterGroup) *gin.RouterGroup {
	userRouteGroup := parent.Group("/user")

	userRouteGroup.POST("/set-jid", AddJID)

	return userRouteGroup
}
