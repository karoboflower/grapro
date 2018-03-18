package middleware

import (
	"gra-pro/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RBAC 基于角色的权限控制访问
func RBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !database.AuthEnforcer.Enforce(c.Param("id"), c.Request.URL.Path, c.Request.Method) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "您无权访问该资源"})
			return
		}
	}
}
