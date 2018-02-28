package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// AuthRequired 处理认证
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Println("AuthRequired report latency : ", latency)
		status := c.Writer.Status()
		log.Println("AuthRequired report status : ", status)
	}
}
