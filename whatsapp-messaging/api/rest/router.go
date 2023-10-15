package rest

import (
	"net/http"
	"whatsapp-messaging/api/rest/auth"
	"whatsapp-messaging/api/rest/login"
	"whatsapp-messaging/api/rest/user"

	"github.com/gin-gonic/gin"
)

func AddRoutes(parent *gin.RouterGroup) {
	api := parent.Group("/api")

	userRouteGroup:= user.AddRoutes(api)

	authenticated := userRouteGroup.Group("/:user_id", APIHandler())
	login.AddRoutes(authenticated)
}

func APIHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, err := auth.Authenticate(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
