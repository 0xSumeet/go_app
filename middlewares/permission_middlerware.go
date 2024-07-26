package middlewares

import (
	"net/http"

	"github.com/0xSumeet/go_app/database"
	"github.com/0xSumeet/go_app/models"
	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(requiredRole, requiredAccess string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		permission := models.Permission{
			Role:   role,
			Module: c.Request.URL.Path,
			Access: requiredAccess,
		}
		if !permission.CheckPermission(database.DB) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
