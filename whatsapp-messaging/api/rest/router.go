package rest

import (
	"net/http"
	"whatsapp-messaging/api/rest/auth"
	"whatsapp-messaging/api/rest/login"

	"github.com/gin-gonic/gin"
)

func AddRoutes(parent *gin.RouterGroup) {
	parent.Group("/api", APIHandler())

	login.AddRoutes(parent)
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
