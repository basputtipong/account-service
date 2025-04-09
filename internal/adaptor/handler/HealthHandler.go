package handler

import (
	"account-service/infrastructure"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HealthHandle(c *gin.Context) {
	err := infrastructure.PingAllDb()
	if err != nil {
		res := map[string]string{
			"status": "unhealthy",
			"msg":    fmt.Sprintf("unhealthy get error: %s", err.Error()),
		}
		c.JSON(500, res)
		return
	}
	c.JSON(200, map[string]string{"status": "healthy"})
}
