package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) Render(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
